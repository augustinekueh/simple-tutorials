package main

import (
	"fmt"
)


func convertToBytes(x []byte){
	for _, v:= range x{
		fmt.Println(v)
	} 
	//return
}

func convertToString(y []byte){
	for _, s := range y{
		fmt.Println(string(s))
	}
	//return
}

func main(){
	//First example
	buf_alp:=[]byte("ABC")
	convertToBytes(buf_alp)

	//Second Example
	buf_num:=[]byte{65, 66, 67, 226, 78, 172}
	convertToString(buf_num)
	//fmt.Println(string(buf_num))

	//yolo
}