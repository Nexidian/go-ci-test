package main

import "fmt"

func sum(x int, y int) int {
	return x + y

}

func main() {
	fmt.Print(sum("string", 2))
}