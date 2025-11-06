package goroutine

import "fmt"

func HelloName(name string) string {
	return "Hallo " + name
}

func DisplayNumber(num int) {
	fmt.Println("Display", num)
}
