package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// var msg = flag.String("msg", "デフォルト値", "説明")
var n = flag.Bool("n", false, "line number option")

// func init() {
// 	//ポインタを指定して設定を予約
// 	flag.IntVar(&n, "n", 1, "回数")
// }

func main() {
	flag.Parse()
	// fmt.Println(strings.Repeat(*msg, n))

	//引数の数だけループする
	for _, arg := range flag.Args() {
		// fmt.Println(arg)
		err := readFile(arg)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(*n)
}

func readFile(fn string) error {
	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	return nil
}
