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
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		if len(nombres) < 2 {
			return nil, errors.New("la pile n'est pas correcte")
		}
		n1 := nombres[len(nombres)-2]
		n2 := nombres[len(nombres)-1]
		var n = 0
		if s == "+" {
			n = n1 + n2
		} else if s == "-" {
			n = n1 - n2
		} else if s == "*" {
			n = n1 * n2
		} else if s == "/" {
			if n2 == 0 {
				return nil, errors.New("division par zero")
			}
			n = n1 / n2
		} else if s == "%" {
			if n2 == 0 {
				return nil, errors.New("division par zero")
			}
			n = n1 % n2
		}
		nombres = nombres[:len(nombres)-2]
		nombres = append(nombres, n)
	} else {
		return nil, errors.New("opérateur invalide : " + s)
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
				return 0, err
			}
		}
	}

	fmt.Println(nombres)

	if len(nombres) != 1 {
		return 0, errors.New("la pile n'est pas correcte")
	} else {
		return nombres[0], nil
	}
}
