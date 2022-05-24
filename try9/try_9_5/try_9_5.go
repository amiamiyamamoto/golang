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

	// 標準入力に一行受け取る
	ch := input(os.Stdin)
	// 制限時間を設定
	tm := time.After(5 * time.Second)

L:
	for {
		select {
		case m := <-ch:

			fmt.Print(">")
			// 受け取った単語が単語リスト内に含まれているかチェック
			if checkWord(m) {
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
func wordlistSearch(w string, l []string) bool {
	for _, word := range l {
		if word == w {
			return true
		}
	}
	return false
}

// 入力された単語が正解かどうかを判定する関数
func checkWord(w string) bool {
	//単語リストに含まれていなければ不正解
	if !wordlistSearch(w, words) {
		return false
	}
	//入力済み単語リストに含まれている場合不正解
	if wordlistSearch(w, inw) {
		return false
	}
	//正解の場合、inwに単語を追加する
	inw = append(inw, w)
	return true
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

// 正解した単語を収容するスライス
var inw = []string{}
