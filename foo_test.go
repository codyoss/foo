package foo

import "testing"

func TestCombine(t *testing.T) {
	type args struct {
		is []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"one int", args{[]int{1}}, "1 "},
		{"two ints", args{[]int{1, 2}}, "1 2 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combine(tt.args.is...); got != tt.want {
				t.Errorf("Combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
