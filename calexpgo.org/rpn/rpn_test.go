package rpn

import (
	"errors"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		param []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]string{"1", "2", "+"}}, []int{3}},
		{"test2", args{[]string{"4", "5", "+", "2", "-"}}, []int{7}},
		{"test3", args{[]string{"1", "2", "+", ";", "3", "4", "*"}}, []int{3, 12}},
		{"dup", args{[]string{"4", "dup", "+"}}, []int{8}},
		{"DUP", args{[]string{"1", "DUP", "+"}}, []int{2}},
		{"test4", args{[]string{"1", "2", "+", ";", "5", "6", "*"}}, []int{3, 30}},
		{"test5", args{[]string{"1", "2", "+", ";", ":", "NIP", "swap", "drop", ";", "7", "9", "NIP"}}, []int{3, 9}},
		{"test6", args{[]string{":", "CARREE", "7", "dup", "*", ";", "CARREE"}}, []int{49}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Run(tt.args.param); !reflect.DeepEqual(got, tt.want) || err != nil {
				t.Errorf("executeOperation() = (%v,%v) want %v", got, err, tt.want)
			}
		})
	}
}

func TestRunError(t *testing.T) {
	type args struct {
		param []string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"test1", args{[]string{"1", "2", "?"}}, errors.New("opérateur invalide : ?")},
		{"test2", args{[]string{"1", "2", "3", "+"}}, errors.New("la pile n'est pas correcte")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Run(tt.args.param); !reflect.DeepEqual(got, []int{}) || !reflect.DeepEqual(err, tt.want) {
				t.Errorf("executeOperation() = (%v, %v) want %v", got, err, tt.want)
			}
		})
	}
}

func Test_executeOperation(t *testing.T) {
	type args struct {
		s       string
		nombres []int
	}
	type res struct {
		s   []int
		err error
	}
	tests := []struct {
		name string
		args args
		want res
	}{
		{"test_plus", args{"+", []int{1, 2}}, res{[]int{3}, nil}},
		{"test_moins", args{"-", []int{7, 2}}, res{[]int{5}, nil}},
		{"test_fois", args{"*", []int{7, 4}}, res{[]int{28}, nil}},
		{"test_div", args{"/", []int{6, 3}}, res{[]int{2}, nil}},
		{"test_reste_div", args{"%", []int{7, 3}}, res{[]int{1}, nil}},
		{"test_plus2", args{"+", []int{1, 2, 3, 4, 5, 6}}, res{[]int{1, 2, 3, 4, 11}, nil}},
		{"test_moins2", args{"-", []int{1, 2, 3, 4, 7, 4}}, res{[]int{1, 2, 3, 4, 3}, nil}},
		{"test_dup", args{"dup", []int{8}}, res{[]int{8, 8}, nil}},
		{"test_drop", args{"drop", []int{4, 5}}, res{[]int{4}, nil}},
		{"test_swap", args{"swap", []int{4, 5}}, res{[]int{5, 4}, nil}},
		{"test_over", args{"over", []int{4, 5}}, res{[]int{4, 5, 4}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stack = stack{nombres: tt.args.nombres}
			if err := stack.executeOperation(tt.args.s); !reflect.DeepEqual(stack.nombres, tt.want.s) ||
				err != nil {
				t.Errorf("executeOperation() = %v, want %v", stack.nombres, tt.want)
			}
		})
	}
}

func Test_executeOperationError(t *testing.T) {
	type args struct {
		s       string
		nombres []int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"test1", args{"_", []int{7, 2}}, errors.New("opérateur invalide : _")},
		{"test2_un_seul_operateur", args{"+", []int{2}}, errors.New("la pile n'est pas correcte")},
		{"test3_aucun_operateur", args{"+", []int{}}, errors.New("la pile n'est pas correcte")},
		{"test4_division_par_zero", args{"/", []int{8, 0}}, errors.New("division par zero")},
		{"test5_reste_par_zero", args{"%", []int{7, 0}}, errors.New("division par zero")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stack = stack{nombres: tt.args.nombres}
			if err := stack.executeOperation(tt.args.s); !reflect.DeepEqual(err.Error(), tt.want.Error()) {
				t.Errorf("executeOperation() = (%v, '%v'), want '%v'", tt.args.s, err, tt.want)
			}
		})
	}
}

func Test_isInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"chiffres", args{"123"}, true},
		{"lettres", args{"abc"}, false},
		{"chiffres_lettres", args{"abc123"}, false},
		{"chiffres_lettres2", args{"123abc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInt(tt.args.s); got != tt.want {
				t.Errorf("isInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
