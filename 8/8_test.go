package main

import (
	"testing"
)

func Test_setBitI64(t *testing.T) {
	type args struct {
		a     int64
		n     int
		isSet bool
		want  int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set bit value from 0 to 0",
			args: args{
				a:     122,
				n:     0,
				isSet: false,
				want:  122,
			},
		},
		{
			name: "set bit value from 0 to 1",
			args: args{
				a:     122,
				n:     0,
				isSet: true,
				want:  123,
			},
		},
		{
			name: "set bit value from 1 to 0",
			args: args{
				a:     122,
				n:     6,
				isSet: false,
				want:  58,
			},
		},
		{
			name: "set bit value from 1 to 1",
			args: args{
				a:     122,
				n:     6,
				isSet: true,
				want:  122,
			},
		},
		{
			name: "set sign bit from 0 to 1",
			args: args{
				a:     1,
				n:     63,
				isSet: true,
				want:  -0x7F_FF_FF_FF_FF_FF_FF_FF,
			},
		},
		{
			name: "set sign bit from 1 to 0",
			args: args{
				a:     -0x7F_FF_FF_FF_FF_FF_FF_FF,
				n:     63,
				isSet: false,
				want:  1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setBitI64(&(tt.args.a), tt.args.n, tt.args.isSet)
			if tt.args.a != tt.args.want {
				t.Errorf("setBitI64() = %d, want %d", tt.args.a, tt.args.want)
			}
		})
	}
}
