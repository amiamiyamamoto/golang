package main

import "fmt"

func main() {
	fmt.Println(IsEven(5))

}

func IsEven(n int) bool {
	return n%2 == 0
}
