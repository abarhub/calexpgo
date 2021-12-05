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
		want int
	}{
		{"test1", args{[]string{"1", "2", "+"}}, 3},
		{"test1", args{[]string{"4", "5", "+", "2", "-"}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Run(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeOperation() = %v, want %v", got, tt.want)
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
			if got, err := Run(tt.args.param); !reflect.DeepEqual(got, 0) || !reflect.DeepEqual(err, tt.want) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := executeOperation(tt.args.s, tt.args.nombres); !reflect.DeepEqual(got, tt.want.s) {
				t.Errorf("executeOperation() = %v, want %v", got, tt.want)
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
			if got, err := executeOperation(tt.args.s, tt.args.nombres); got != nil || !reflect.DeepEqual(err.Error(), tt.want.Error()) {
				t.Errorf("executeOperation() = (%v, '%v'), want '%v'", got, err, tt.want)
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
