package main

import (
	"fmt"
	"github.com/firmanjs/go-example-import-with-module/helpers"
)

func main(){
	test := helpers.hello("adam")

	fmt.Println(test)
}