package main
 import ("fmt"
		"strconv"
		"strings"
)

 func main(){
	// value:=hexToBinary("3002300230752503","Big endian")
	// value:=hexToBinary("0000FFFFFFFFFF00")
	// decimalValue, _:= strconv.ParseUint("000000000000000F", 16, 64)

	// array:=strings.Split("0000000000000FFF","")
	// var count int=0
    // for _,ch:=range array{
	// 	if(ch=="F"){
	// 		count++;
	// 	}
	// }

   fmt.Println(count)
	// fmt.Println(len(value),"-=>",value)
 }

 func hexToBinary(hexStr string,endiness string) string {
	decimalValue, _:= strconv.ParseUint(hexStr, 16, 64)
	
	binaryStr := fmt.Sprintf("%b", decimalValue)
   
    for len(binaryStr)!=64{
		binaryStr ="0"+binaryStr
	}
	if endiness=="Big endian"{
		return binaryStr
	}else if endiness=="Little endian"{
		arrayofstr:=strings.Split(binaryStr,"")
		newbinarydata:=""
		str:=""
		for _,val:= range arrayofstr{
			str=str+val
			if len(str)%8 == 0{
				newbinarydata=str+newbinarydata
				str=""
			}
		}
		fmt.Println(binaryStr,"=====",newbinarydata)
		return "--------"
	}else{
		return ""
	}
}

// func hexToBinary(maskbitshex) int{

// return 1
// }