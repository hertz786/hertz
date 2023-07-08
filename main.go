package main

import (
	"fmt"

	controller "github.com/hertz786/hertz/controllers"
	mainController "github.com/hertz786/hertz/controllers/main"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf(" ali hamza branch hello, world\n")
	fmt.Printf("hello, world again\n")
	n := 10
	a, b := mainController.Maincont(n)
	fmt.Println("this is value\n", a, b)
	controller.Mainconfunc()
}
