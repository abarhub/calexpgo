package rpn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type stacki interface {
	errorStackSize(nombres []int, minimalSize int) error
	push(valeur int)
	pop() (int, error)
	len() int
	executeOperation(s string) error
}

type stack struct {
	nombres []int
}

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

func (s *stack) errorStackSize(minimalSize int) error {
	if len(s.nombres) < minimalSize {
		return fmt.Errorf("la pile n'est pas correcte")
	}
	return nil
}

func (s *stack) push(valeur int) {
	s.nombres = append(s.nombres, valeur)
}

func (s *stack) pop() (int, error) {
	err := s.errorStackSize(1)
	if err != nil {
		return 0, err
	}
	n := s.nombres[len(s.nombres)-1]
	s.nombres = s.nombres[:len(s.nombres)-1]
	return n, nil
}

func (s *stack) len() int {
	return len(s.nombres)
}

func (stack *stack) executeOperation(s string) error {
	var err error
	if s == "+" || s == "-" || s == "*" || s == "/" || s == "%" {
		err = stack.errorStackSize(2)
		if err != nil {
			return err
		}
		var n1, n2 int
		n2, err = stack.pop()
		if err != nil {
			return err
		}
		n1, err = stack.pop()
		if err != nil {
			return err
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
				return fmt.Errorf("division par zero")
			}
			n = n1 / n2
		} else if s == "%" {
			if n2 == 0 {
				return fmt.Errorf("division par zero")
			}
			n = n1 % n2
		}
		stack.push(n)
	} else if strings.ToLower(s) == "dup" {
		err = stack.errorStackSize(1)
		if err != nil {
			return err
		}
		n := stack.nombres[stack.len()-1]
		stack.push(n)
	} else if strings.ToLower(s) == "drop" {
		err = stack.errorStackSize(1)
		if err != nil {
			return err
		}
		_, err = stack.pop()
		if err != nil {
			return err
		}
	} else if strings.ToLower(s) == "swap" {
		err = stack.errorStackSize(2)
		if err != nil {
			return err
		}
		var n1, n2 int
		n2, err = stack.pop()
		if err != nil {
			return err
		}
		n1, err = stack.pop()
		if err != nil {
			return err
		}
		stack.push(n2)
		stack.push(n1)
	} else if strings.ToLower(s) == "over" {
		err = stack.errorStackSize(2)
		if err != nil {
			return err
		}
		n1 := stack.nombres[stack.len()-2]
		stack.push(n1)
	} else {
		return fmt.Errorf("opérateur invalide : " + s)
	}
	return nil
}

func Run(param []string) ([]int, error) {
	var stack = stack{nombres: make([]int, 0, 100)}

	var resultat = []int{}
	var programmes = make(map[string][]string)

	for i := 0; i < len(param); i++ {
		s := param[i]
		if isInt(s) {
			v, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				stack.push(int(v))
			} else {
				return []int{}, fmt.Errorf("erreur pour parser le nombre: %v", err)
			}
		} else if s == ";" {
			if stack.len() > 0 {
				var n int
				var err error
				n, err = stack.pop()
				if err != nil {
					return []int{}, err
				}
				resultat = append(resultat, n)
			}
		} else if s == ":" {
			i++
			nom_programme := param[i]
			i++
			var instr []string
			for ; i < len(param); i++ {
				s = param[i]
				if s == ";" {
					break
				} else {
					instr = append(instr, s)
				}
			}
			programmes[nom_programme] = instr
		} else if v, found := programmes[s]; found {
			var instr = v
			for j := 0; j < len(instr); j++ {
				var err error
				s := instr[j]
				err = stack.executeOperation(s)
				if err != nil {
					return []int{}, err
				}
			}
		} else {
			var err error
			err = stack.executeOperation(s)
			if err != nil {
				return []int{}, err
			}
		}
	}

	if len(programmes) == 0 || stack.len() == 1 {
		if stack.len() != 1 {
			return []int{}, fmt.Errorf("la pile n'est pas correcte")
		} else {
			resultat = append(resultat, stack.nombres[0])
			return resultat, nil
		}
	} else {
		return resultat, nil
	}
}
