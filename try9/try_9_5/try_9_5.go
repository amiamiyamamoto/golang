package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// 標準出力に英単語を出す
	for _, w := range words {
		fmt.Println(w)
	}

	//TODO 一分経ったら標準出力に「終了」の文字を表示させる
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("終了")
	// }()

	// 標準入力に一行受け取る
	ch := input(os.Stdin)

	tm := time.After(5 * time.Second)

L:
	for {
		select {
		case m := <-ch:

			fmt.Print(">")
			// 受け取った単語が単語リスト内に含まれているかチェック
			if wordlistSearch(m) {
				fmt.Println("OK")
				point++
			}

			// fmt.Println(point)
		case <-tm:
			fmt.Println("\ntime out!")
			fmt.Println(point)
			break L
		}

	}

	// for {
	// 	fmt.Print(">")
	// 	// 受け取った単語が単語リスト内に含まれているかチェック
	// 	if wordlistSearch(<-ch) {
	// 		fmt.Println("OK")
	// 		point++
	// 	}
	// 	fmt.Println(point)
	// }
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
var point int

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

// TODO: 同じ単語を2回入力しても点数は最初の1点のみにする
