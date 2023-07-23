package main

import (
	"fmt"
)

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func main() {

	fmt.Printf("%t\n", IsDigit('8'))

}
