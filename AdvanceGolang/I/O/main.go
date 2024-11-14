package main

import (
	"fmt"
	"io"
	"strings"
)


func main(){
	read := strings.NewReader(" I Am NurediN Bedru Try To Learn Golang File Managment !!!")

	Data , err := io.ReadAll(read)

	if err != nil{
		panic(err)

	}

	fmt.Println("Read Data Is : ", Data)


}