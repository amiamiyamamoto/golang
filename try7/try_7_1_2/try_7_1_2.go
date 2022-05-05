package main

import (
	"scanner/myscanner"
	"strings"
)

func main() {

	// fmt.Println(myscanner.NewScanner(strings.NewReader("あいうえお")))
	myscanner.NewScanner(strings.NewReader("あいうabc123"))
	// fmt.Printf("%+v\n", *a)
}
