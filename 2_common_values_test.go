
package main

import (
	"reflect"
	"testing"
)

func TestCommonValues(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one number in common",
			args: args{a: []int{100, 200, 300}, b: []int{200, 400, 500}},
			want: []int{200},
		},
		{
			name: "two number in common",
			args: args{a: []int{100, 200, 300, 400}, b: []int{200, 400, 500}},
			want: []int{200, 400},
		},
		{
			name: "no numbers in common",
			args: args{a: []int{100, 200, 300, 400}, b: []int{500, 600}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonValues(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
