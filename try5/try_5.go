package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// var msg = flag.String("msg", "デフォルト値", "説明")
var n = flag.Bool("n", false, "line number option")
var num int64

// func init() {
// 	//ポインタを指定して設定を予約
// 	flag.IntVar(&n, "n", 1, "回数")
// }

func main() {
	flag.Parse()

	//nオプションが付いている場合、numに数値をセットする
	if *n {
		num = 1
	}

	//引数の数だけループする
	for _, arg := range flag.Args() {
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
	var line_num string
	for s.Scan() {
		if num != 0 {
			line_num = strconv.FormatInt(num, 10) + ": "
			num++
		}
		fmt.Println(line_num, s.Text())

	}
	return nil
}
