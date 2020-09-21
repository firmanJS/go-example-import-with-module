package main

import (
	"fmt"
	"github.com/firmanJS/go-example-import-with-module/helpers"
)

func main(){
	test := helpers.hello()

	fmt.Println(test)
}