package main

import (
	"testing"
)

func TestXFormula(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{

		{
			name: "3 should return 3702",
			args: args{n: 3},
			want: 3702,
		},
		{
			name: "4 should return 4936",
			args: args{n: 4},
			want: 4936,
		},
		{
			name: "9 should return 4936",
			args: args{n: 9},
			want: 11106,
		},
		{
			name: "10 should return error",
			args: args{n: 10},
			wantErr : true,
		},
		{
			name: "-1 should return error",
			args: args{n: -1},
			wantErr : true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XFormula(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("XFormula() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("XFormula() got = %v, want %v", got, tt.want)
			}
		})
	}
}