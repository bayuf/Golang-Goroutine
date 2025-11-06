package main

import (
	"fmt"
	"runtime"
)

func print(till int, name string) {
	for i := 1; i <= till; i++ {
		fmt.Println(i, "Hallo", name)
	}
}

func main() {
	runtime.GOMAXPROCS(5)

	go print(10, "Bayu")
	go print(10, "Silfi")
	print(10, "Dunia")

	var input string
	fmt.Scanln(&input)
}
