package main

import (
	"calexpgo/calexpgo.org/rpn"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	res, err := rpn.Run(argsWithoutProg)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		if res == nil {
			fmt.Println("error: no result")
		} else {
			for _, n := range res {
				fmt.Println(n)
			}
		}
	}

}
