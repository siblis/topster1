package cnsl

import "fmt"

//Read - reads from console and prints name
func Read() {
	var name string
	fmt.Print("Your name? ")
	fmt.Scanln(&name)
	fmt.Println(name)

}
