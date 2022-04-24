package main

import "fmt"

type Stringer interface {
	String() string
}

type StringerInt int

func (s StringerInt) Stringer() string {
	return string(s)
}

func main() {
	var s StringerInt = 10
	fmt.Println(s.Stringer())
}
