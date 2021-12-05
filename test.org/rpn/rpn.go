package rpn

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func executeOperation(s string, nombres []int) ([]int, error) {
	if s == "+" {
		n1 := nombres[len(nombres)-1]
		n2 := nombres[len(nombres)-2]
		n := n1 + n2
		nombres = nombres[:len(nombres)-2]
		nombres = append(nombres, n)
	} else if s == "-" {
		n1 := nombres[len(nombres)-1]
		n2 := nombres[len(nombres)-2]
		n := n2 - n1
		nombres = nombres[:len(nombres)-2]
		nombres = append(nombres, n)
	}
	return nombres, nil
}

func Run(param []string) (int, error) {
	var nombres = make([]int, 0, 100)

	for i := 0; i < len(param); i++ {
		s := param[i]
		if isInt(s) {
			v, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				nombres = append(nombres, int(v))
			}
		} else {
			err := errors.New("")
			nombres, err = executeOperation(s, nombres)
			if err != nil {
				fmt.Errorf("erreur: %s", err)
				break
			}
		}
	}

	fmt.Println(nombres)

	return nombres[0], nil
}
