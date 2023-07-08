package controller

import (
	"fmt"
)

func Maincont(n int) (int, int) {
	fmt.Printf("this is function maincont from /controller/main/firstController.go")
	m := n
	g := m + 1
	return m, g
}
