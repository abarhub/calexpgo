package rpn

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Run(tt.args.param)
			if got, _ := Run(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeOperation(t *testing.T) {
	type args struct {
		s       string
		nombres []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{"+", []int{1, 2}}, []int{3}},
		{"test2", args{"-", []int{7, 2}}, []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := executeOperation(tt.args.s, tt.args.nombres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeOperation() = %v, want %v", got, tt.want)
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
