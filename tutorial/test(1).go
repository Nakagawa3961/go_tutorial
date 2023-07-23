package main

import (
	"fmt"
	"os"
)

func IsDigit(c int) bool {
	return c >= '0' && c <= '9'
}

func main() {

	// Args[1:]でコマンドライン引数を取得
	args := os.Args[1:]

	for _, arg := range args {
		for _, char := range arg {
			if IsDigit(int(char)) {
				fmt.Printf("%c True\n", char)
			} else {
				fmt.Printf("%c Flse\n", char)
			}
		}
	}
}
