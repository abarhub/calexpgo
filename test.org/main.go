package main

import (
	"calexpgo/test.org/rpn"
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)

	res, err := rpn.Run(argsWithoutProg)

	if err != nil {
		fmt.Errorf("error: %s", err)
	} else {
		fmt.Println(res)
	}

}
