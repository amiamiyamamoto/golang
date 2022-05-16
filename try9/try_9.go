package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}

func input(r io.Reader) <-chan string {
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			//TODO: チャネルに読み込んだ文字列を送る
		}
		//TODO: チャネルを閉じる
	}()
	//TODO: チャネルを返す
}
