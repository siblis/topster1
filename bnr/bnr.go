package bnr

import "fmt"

const b0010 = 0x1 << 2

//Tostring - convert digits to string in binary presentation
func Tostring(b int) string {
	return fmt.Sprintf("%04b", b)
}

//Println - print binary
func Println(b int) {
	fmt.Println(Tostring(b))
}

//Bxxx - binary digits
const (
	B0000 = 0x0
	B0001 = 0x1
	B0010 = 0x1 << 1
	B0011 = B0010 | B0001
	B0100 = 0x1 << 2
	B0101 = 0x5
	B0110 = 0x6
	B0111 = 0x7
	B1000 = 0x8
	B1001 = 0x8 | B0001
	B1010 = 0x8 | B0010
)
