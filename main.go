package main
import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io"
	"encoding/csv"
	"os"
	"strconv"
)

func main(){
  token := Gettoken()
  configid:="C20BF6BDB3E6A761D157A215A6585F2C148EF15D0C0F5FDC0306F6FD00D6ED6C"
  Data := GetPidMapping(configid,token)
// //   fmt.Println(Data)
// for i,value :=range Data.Data{
// 	fmt.Println(i,value.Name)
// }


data, err := readCSVFile("../../Downloads/laf_297078_11_10_2024.csv")
if err != nil {
	fmt.Println("Error reading file:", err)
	return
}
reader, err := parseCSV(data)
if err != nil {
	fmt.Println("Error creating CSV reader:", err)
	return
}
canData,_ := processCSV(reader);

// fmt.Println(canData)
for key,_ := range canData{
	pidmap := getMappingbits(key,Data)
	processData(key,canData[key],pidmap)
	fmt.Println(key,canData[key],pidmap)
	break
	
}

}
func getMappingbits(pid string,Data pidMapping) []pidData {
pidMapArr :=[]pidData{}

for _, value := range Data.Data {
	id := strconv.Itoa(value.PidCode) // PidCode to string
	if pid == id {
		pidMapArr = append(pidMapArr, value)
	}
}

return pidMapArr
}

func processData(pidId string,canData []string,pidmap []pidData){

fmt.Println(pidId,canData,pidmap)

}















func Gettoken() string {

    postBody, _ := json.Marshal(map[string]map[string]string{
        "user": {
            "type":     "apiuser",
            // "username": "lg_immob.api",
            "username": "bsmosfet.api",
            "password": "intellicar@123",
        }})

    responseBody := bytes.NewBuffer(postBody)
    resp, err := http.Post("https://apiplatform.intellicar.in/gettoken", "application/json", responseBody)

    if err != nil {
        log.Fatalf("An Error Occured %v", err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    tokenResp := struct {
        Status string `json:"status"`
        Data   struct {
            Token string `json:"token"`
        } `json:"data"`
        Userinfo struct {
            Userid   int    `json:"userid"`
            Typeid   int    `json:"typeid"`
            Username string `json:"username"`
        } `json:"userinfo"`
        Err string `json:"err"`
        Msg string `json:"msg"`
    }{}
    if err := json.Unmarshal(body, &tokenResp); err != nil {
        log.Fatal("Token parsing error")
        fmt.Print(err)
    }

    Token := tokenResp.Data.Token
      
    fmt.Println(Token)
    return Token
}


func readCSVFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    data, err := io.ReadAll(f)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error) {
    reader := csv.NewReader(bytes.NewReader(data))
    return reader, nil
}


func GetPidMapping(configid string,token string) pidMapping{
	postBody,_:=json.Marshal(map[string]string{
		"configid": configid,
		"deviceid":"",
		"publickey":"",
		"token":token,
	})
	
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://laf.intellicar.in/api/can/getlafpidmappingswtoken","application/json",responseBody)
	
	if err !=nil{
		log.Fatalf("An Error Occured %v",err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}
	
	var pidMappingResponse pidMapping
	if err := json.Unmarshal(body,&pidMappingResponse); err !=nil{
		log.Fatal("pidmapping parsing error")
		fmt.Print(err)
	}
	
	fmt.Println(pidMappingResponse.Status)
	return pidMappingResponse
}

type pidData struct {
	Name              string `json:"name"`
	PidCode           int    `json:"pidcode"`
	ShiftBits         int    `json:"shiftbits"`
	MaskBitsHex       string `json:"maskbitshex"`
	Multiplier        int    `json:"multiplier"`
	Divisor           int    `json:"divisor"`
	Offset            int    `json:"offset"`
	Endian            int    `json:"endian"`
	IsCustomConversion int   `json:"iscustomconversion"`
}

type pidMapping struct {
	Status string   `json:"status"`
	Data   []pidData `json:"data"`
}
// "name": "afe_acquisition_chip_malfunction",
//             "pidcode": 3018,
//             "shiftbits": 40,
//             "maskbitshex": "1",
//             "multiplier": 1,
//             "divisor": 1,
//             "offset": 0,
//             "endian": 0,
//             "iscustomconversion": 0


func ProcessRawData(Data pidMapping,csvdata *csv.Reader){

	for{
		record,err :=csvdata.Read()
		if err == io.EOF {
			break
		} else if err != nil{
			fmt.Println("Error reading CSV data:",err)
			break
		}
		fmt.Println(record)
	}
   
}

func processCSV(reader *csv.Reader) (map[string][]string,error) {
	pidvaluemap := make(map[string][]string);
	
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return nil,err
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break 
			} else if err != nil {
				fmt.Println("Error reading CSV data:", err)
				return nil,err
			}
			
		for i, value := range record {
			// fmt.Println(value)
			pidvaluemap[header[i]] = append(pidvaluemap[header[i]],value)
		}
	}

	fmt.Println("---------------------")
return pidvaluemap,nil
}