package main

import (
	"fmt"

	controller "github.com/hertz786/hertz/controllers"
	mainController "github.com/hertz786/hertz/controllers/main"
)

func main() {
	fmt.Printf("hello, world")
	fmt.Printf("hello, world again")
	n := 10
	a, b := mainController.Maincont(n)
	fmt.Println("this is value", a, b)
	controller.Mainconfunc()
}
