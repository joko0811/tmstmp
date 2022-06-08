package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joko0811/timestamp/data"
)


func main() {
	wd, err:=os.Getwd()
	if(err!=nil){
		log.Println("err",err)
	}
	fmt.Println(wd+data.config.Filename)
}
