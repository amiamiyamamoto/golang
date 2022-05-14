package main

import "fmt"

func main() {
	fmt.Println(isEven(5))

}

func isEven(n int) bool {
	return n%2 == 0
}
