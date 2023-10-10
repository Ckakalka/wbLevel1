package main

import (
	"encoding/binary"
	"testing"
)

func Test_setBitI64(t *testing.T) {
	type args struct {
		a     int64
		n     int
		isSet bool
		want  int64
	}
	testsLittleEndian := []struct {
		name string
		args args
	}{
		{
			name: "LE set bit value from 0 to 0",
			args: args{
				a:     122,
				n:     0,
				isSet: false,
				want:  122,
			},
		},
		{
			name: "LE set bit value from 0 to 1",
			args: args{
				a:     122,
				n:     0,
				isSet: true,
				want:  123,
			},
		},
		{
			name: "LE set bit value from 1 to 0",
			args: args{
				a:     122,
				n:     6,
				isSet: false,
				want:  58,
			},
		},
		{
			name: "LE set bit value from 1 to 1",
			args: args{
				a:     122,
				n:     6,
				isSet: true,
				want:  122,
			},
		},
		{
			name: "LE set sign bit from 0 to 1",
			args: args{
				a:     1,
				n:     63,
				isSet: true,
				want:  -0x7F_FF_FF_FF_FF_FF_FF_FF,
			},
		},
		{
			name: "LE set sign bit from 1 to 0",
			args: args{
				a:     -0x7F_FF_FF_FF_FF_FF_FF_FF,
				n:     63,
				isSet: false,
				want:  1,
			},
		},
	}
	// the value of n is converted as if it is big endian
	testsBigEndian := []struct {
		name string
		args args
	}{
		{
			name: "BE set bit value from 0 to 0",
			args: args{
				a:     122,
				n:     56,
				isSet: false,
				want:  122,
			},
		},
		{
			name: "BE set bit value from 0 to 1",
			args: args{
				a:     122,
				n:     56,
				isSet: true,
				want:  123,
			},
		},
		{
			name: "BE set bit value from 1 to 0",
			args: args{
				a:     122,
				n:     62,
				isSet: false,
				want:  58,
			},
		},
		{
			name: "BE set bit value from 1 to 1",
			args: args{
				a:     122,
				n:     62,
				isSet: true,
				want:  122,
			},
		},
		{
			name: "BE set sign bit from 0 to 1",
			args: args{
				a:     1,
				n:     7,
				isSet: true,
				want:  -0x7F_FF_FF_FF_FF_FF_FF_FF,
			},
		},
		{
			name: "BE set sign bit from 1 to 0",
			args: args{
				a:     -0x7F_FF_FF_FF_FF_FF_FF_FF,
				n:     7,
				isSet: false,
				want:  1,
			},
		},
	}

	for _, tt := range testsLittleEndian {
		nativeEndian = binary.LittleEndian
		t.Run(tt.name, func(t *testing.T) {
			setBitI64(&(tt.args.a), tt.args.n, tt.args.isSet)
			if tt.args.a != tt.args.want {
				t.Errorf("setBitI64() = %d, want %d", tt.args.a, tt.args.want)
			}
		})
	}
	for _, tt := range testsBigEndian {
		// this is not real big endian testing
		nativeEndian = binary.BigEndian
		t.Run(tt.name, func(t *testing.T) {
			setBitI64(&(tt.args.a), tt.args.n, tt.args.isSet)
			if tt.args.a != tt.args.want {
				t.Errorf("setBitI64() = %d, want %d", tt.args.a, tt.args.want)
			}
		})
	}

}
