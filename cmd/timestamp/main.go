package main

import (
	"fmt"
	"os"
	"log"
)


func main() {
	wd, err:=os.Getwd()
	if(err!=nil){
		log.Println("err",err)
	}
	fmt.Println(wd)
}
