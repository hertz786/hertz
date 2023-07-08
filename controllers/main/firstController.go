package controller

import (
	"fmt"
)

func Maincont(n int) (int, int) {
	fmt.Printf("\nthis is function maincont from /controller/main/firstController.go\n")
	m := n
	g := m + 1
	return m, g
}
