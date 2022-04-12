package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	//ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}
