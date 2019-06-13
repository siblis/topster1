package main

import (
	"fmt"
	"strconv"
	"topster1/bnr"
)

func main() {
	const umask=0022
	const executeaccess=bnr.B0001
	const writeaccess=bnr.B0010
	const readaccess=bnr.B0100
	fmt.Printf("%03o\n",writeaccess|readaccess)

	  
	
	mask := bnr.B0110 ^ bnr.B1001
	bnr.Println(mask)
	bnr.Println(bnr.B0001 << 2)
	bnr.Println(bnr.B0001 << 3)


	fmt.Println(strconv.FormatInt(bnr.B0010, 2))
	bnr.Println(bnr.B0010)
	fmt.Println(bnr.Tostring(bnr.B0011))
	{var a,b int
	b=1
	a=b
	a=2
	println("&a=",&a,"&b=",&b)
	println("a=",a,"b=",b)

}
	var a,b *int
	a=new(int)
	b=new(int)
	*b=1
	*a=*b
	*a=2
	println("a=",a,"b=",b)
	println("*a=",*a,"*b=",*b)

	*b=3
	a=b
	*a=2
	println("a=",a,"b=",b)
	println("*a=",*a,"*b=",*b)



}
