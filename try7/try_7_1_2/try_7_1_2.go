package main

import (
	"fmt"
	"os"
	"scanner/myscanner"
	"strings"
)

func main() {

	// fmt.Println(myscanner.NewScanner(strings.NewReader("あいうえお")))
	// s, _ := myscanner.NewScanner(strings.NewReader("あいうabc123"))
	// fmt.Printf("%+v\n", *a)
	// fmt.Println(string(s.Data))
	readMyScanner()
}

func readMyScanner() {
	scanner := myscanner.NewScanner(strings.NewReader("あいうabc123"))
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// NewScannerではエラーを返さない
	// readメソッドでエラーがあった場合は、Scanner structにエラー値をもたせるconst
	// エラーを保持している場合は、Scanメソッドはfalseを返す
	// Errメソッドを作成して、どうやって返せばいいんだ？
}
