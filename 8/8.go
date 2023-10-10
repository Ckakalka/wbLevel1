package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}

// a in little endian notation
func setBitI64(a *int64, n int, isSet bool) {
	if !(0 <= n && n < 64) {
		return
	}
	if nativeEndian == binary.BigEndian {
		// Bits in little endian
		// (63...56)(55...48)(47...40)(39...32)(31...24)(23...16)(15...8) (7... 0)
		//  |    |   |    |   |    |   |    |   |    |   |    |   |    |   |    |
		// (7... 0) (15...8) (23...16)(31...24)(39...32)(47...40)(55...48)(63...56)
		// Bits in big endian
		// Translate n in little endian to the corresponding value of big endian
		n = 64 - 8*(n/8+1) + n%8
	}
	if isSet {
		*a |= (1 << n)
	} else {
		*a &= ^(1 << n)
	}
}

func main() {
	var a int64 = 122
	fmt.Println(a)
	setBitI64(&a, 0, true)
	fmt.Println(a)
}
