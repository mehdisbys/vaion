package main

import (
	"testing"
)

func TestDecimalRepresentation(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no odd digit in decimal representation",
			args: args{n: 222},
			want: true,
		},
		{
			name: "odd digit in decimal representation",
			args: args{n: 232},
			want: false,
		},
		{
			name: "zero",
			args: args{n: 0},
			want: true,
		},
		{
			name: "only odd digits",
			args: args{n: 555},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalRepresentation(tt.args.n); got != tt.want {
				t.Errorf("DecimalRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
