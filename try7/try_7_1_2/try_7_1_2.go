package main

import (
	"fmt"
	"scanner/myscanner"
	"strings"
)

func main() {

	// fmt.Println(myscanner.NewScanner(strings.NewReader("あいうえお")))
	s, _ := myscanner.NewScanner(strings.NewReader("あいうabc123"))
	// fmt.Printf("%+v\n", *a)
	fmt.Println(string(s.Data))
}
