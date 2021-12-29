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

type programi interface {
	addProgram(name string, instr []string)
	findProgram(name string) ([]string, bool)
}

type stack struct {
	nombres []int
}

type program struct {
	programmes map[string][]string
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

func (program *program) addProgram(name string, instr []string) {
	name = strings.ToLower(name)
	program.programmes[name] = instr
}

func (program program) findProgram(name string) ([]string, bool) {
	name = strings.ToLower(name)
	instr, found := program.programmes[name]
	return instr, found
}

func execute(list3 []string, stack stack, programmes program, resultat []int) ([]int, error) {
	for i := 0; i < len(list3); i++ {
		s := list3[i]
		if s == ";" {
			return []int{}, fmt.Errorf("erreur pour parser les instructions: point virgule au mauvais endroit")
		}
		if isInt(s) {
			v, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				stack.push(int(v))
			} else {
				return []int{}, fmt.Errorf("erreur pour parser le nombre: %v", err)
			}
		} else if s == ":" {
			i++
			nom_programme := list3[i]
			i++
			var instr []string
			for ; i < len(list3); i++ {
				s = list3[i]
				instr = append(instr, s)
			}
			programmes.addProgram(nom_programme, instr)
		} else if instr, found := programmes.findProgram(s); found {
			var err error
			resultat, err = execute(instr, stack, programmes, resultat)
			if err != nil {
				return []int{}, err
			}
		} else {
			var err error
			err = stack.executeOperation(s)
			if err != nil {
				return []int{}, err
			}
		}
	}

	if len(programmes.programmes) == 0 || stack.len() == 1 {
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

func Run(param []string) ([]int, error) {
	var stack = stack{nombres: make([]int, 0, 100)}

	var resultat = []int{}
	var programmes = program{programmes: make(map[string][]string)}

	var list = [][]string{}
	var list2 = []string{}

	for i := 0; i < len(param); i++ {
		s := param[i]
		if s == ";" {
			list = append(list, list2)
			list2 = []string{}
		} else {
			list2 = append(list2, s)
		}
	}

	if len(list2) > 0 {
		list = append(list, list2)
	}

	for j := 0; j < len(list); j++ {
		list3 := list[j]
		var err error = nil
		resultat, err = execute(list3, stack, programmes, resultat)
		if err != nil {
			return []int{}, err
		}
	}

	return resultat, nil
}
