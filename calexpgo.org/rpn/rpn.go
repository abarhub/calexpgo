package rpn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isInt(s string) bool {
	if strings.Trim(s, " ") == "" {
		return false
	}
	if len(s) >= 2 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func errorStackSize(nombres []int, minimalSize int) error {
	if len(nombres) < minimalSize {
		return fmt.Errorf("la pile n'est pas correcte")
	}
	return nil
}

func push(nombres []int, valeur int) []int {
	return append(nombres, valeur)
}

func pop(nombres []int) ([]int, int, error) {
	err := errorStackSize(nombres, 1)
	if err != nil {
		return nil, 0, err
	}
	n := nombres[len(nombres)-1]
	return nombres[:len(nombres)-1], n, nil
}

func executeOperation(s string, nombres []int) ([]int, error) {
	var err error
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		err = errorStackSize(nombres, 2)
		if err != nil {
			return nil, err
		}
		var n1, n2 int
		nombres, n2, err = pop(nombres)
		if err != nil {
			return nil, err
		}
		nombres, n1, err = pop(nombres)
		if err != nil {
			return nil, err
		}
		var n = 0
		if s == "+" {
			n = n1 + n2
		} else if s == "-" {
			n = n1 - n2
		} else if s == "*" {
			n = n1 * n2
		} else if s == "/" {
			if n2 == 0 {
				return nil, fmt.Errorf("division par zero")
			}
			n = n1 / n2
		} else if s == "%" {
			if n2 == 0 {
				return nil, fmt.Errorf("division par zero")
			}
			n = n1 % n2
		}
		nombres = push(nombres, n)
	} else if strings.ToLower(s) == "dup" {
		err = errorStackSize(nombres, 1)
		if err != nil {
			return nil, err
		}
		n := nombres[len(nombres)-1]
		nombres = push(nombres, n)
	} else if strings.ToLower(s) == "drop" {
		err = errorStackSize(nombres, 1)
		if err != nil {
			return nil, err
		}
		nombres, _, err = pop(nombres)
		if err != nil {
			return nil, err
		}
	} else if strings.ToLower(s) == "swap" {
		err = errorStackSize(nombres, 2)
		if err != nil {
			return nil, err
		}
		var n1, n2 int
		nombres, n2, err = pop(nombres)
		if err != nil {
			return nil, err
		}
		nombres, n1, err = pop(nombres)
		if err != nil {
			return nil, err
		}
		nombres = push(nombres, n2)
		nombres = push(nombres, n1)
	} else if strings.ToLower(s) == "over" {
		err = errorStackSize(nombres, 2)
		if err != nil {
			return nil, err
		}
		n1 := nombres[len(nombres)-2]
		nombres = push(nombres, n1)
	} else {
		return nil, fmt.Errorf("opérateur invalide : " + s)
	}
	return nombres, nil
}

func Run(param []string) ([]int, error) {
	var nombres = make([]int, 0, 100)

	var resultat = []int{}

	for i := 0; i < len(param); i++ {
		s := param[i]
		if isInt(s) {
			v, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				nombres = push(nombres, int(v))
			} else {
				return []int{}, fmt.Errorf("erreur pour parser le nombre: %v", err)
			}
		} else if s == ";" {
			if len(nombres) > 0 {
				var n int
				var err error
				nombres, n, err = pop(nombres)
				if err != nil {
					return []int{}, err
				}
				//fmt.Println(n)
				resultat = append(resultat, n)
			}
		} else {
			var err error
			nombres, err = executeOperation(s, nombres)
			if err != nil {
				return []int{}, err
			}
		}
	}

	if len(nombres) != 1 {
		return []int{}, fmt.Errorf("la pile n'est pas correcte")
	} else {
		resultat = append(resultat, nombres[0])
		return resultat, nil
	}
}
