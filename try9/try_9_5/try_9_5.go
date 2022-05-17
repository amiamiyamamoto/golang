package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 標準出力に英単語を出す
	for _, w := range words {
		fmt.Println(w)
	}
	// 標準入力に一行受け取る
	ch := input(os.Stdin)

	for {
		fmt.Print(">")
		// fmt.Print(<-ch)
		// 受け取った単語が単語リスト内に含まれているかチェック
		if wordlistSearch(<-ch) {
			fmt.Println("OK")
			correct++
		}
	}
	// 制限時間内に難問解けたか表示する

}

func input(r io.Reader) <-chan string {
	var ch = make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

// 単語リストに含まれているかチェックする関数
func wordlistSearch(w string) bool {
	for _, word := range words {
		if word == w {
			return true
		}
	}
	return false
}

// 正答数をカウントする
var correct int

// 単語リスト
var words = []string{
	"thirteen",
	"friend",
	"theirs",
	"again",
	"welcome",
	"grandfather",
	"grandmother",
	"drive",
	"letter",
	"classroom",
	"fourteen",
	"year",
	"bad",
	"eat",
	"rain",
	"child",
	"morning",
	"Sunday",
	"play",
	"stop",
}
