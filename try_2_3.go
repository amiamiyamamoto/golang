package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 100 + 200
	fmt.Printf("%v\n", n)
	m := n + 100
	fmt.Printf("%v\n", m)
	msg := "hoge" + "huga"
	fmt.Printf("%v\n", msg)

	omikuji()
}

func omikuji() {
	// サイコロをふる
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) + 1
	fmt.Printf("サイコロの目は%vです！\n", n)
	// if n == 6 {
	// 	fmt.Printf("大吉")
	// 	return
	// }
	// if n >= 4 {
	// 	fmt.Println("中吉")
	// 	return
	// }
	// if n >= 2 {
	// 	fmt.Println("小吉")
	// 	return
	// }
	// fmt.Printf("凶")
	// return
	switch n {
	case 6:
		fmt.Printf("大吉")
	case 5, 4:
		fmt.Printf("中吉")
	case 3, 2:
		fmt.Printf("小吉")
	case 1:
		fmt.Printf("凶")
	}
}
