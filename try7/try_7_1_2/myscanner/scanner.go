package myscanner

import (
	"fmt"
	"io"
)

type Scanner struct {
	rune    rune
	reader  io.Reader
	length  uint // 受け取った文字列長を格納
	current uint // 現在読み込んでいる文字位置を格納
}

// スキャンする文字がまだある場合はtrueを返す
func (s *Scanner) Scan() bool {
	return true
}

// 現在の位置の文字を一つ返す
func (s *Scanner) Text() string {
	// s.text.
	return "a"
}

// io.Readerを受け取り、Scannerのポインタを返す
func NewScanner(r io.Reader) *Scanner {
	var s Scanner
	s.reader = r
	// スライスを確保
	//どれくらいの長さになるか事前に確認したいが一旦20で固定
	// 丁度いい長さのスライスを渡したい
	p := make([]byte, 20)
	_, err := r.Read(p)
	fmt.Printf("%+v\n %+v", p, err)

	for _, byte := range p {
		// fmt.Println(string(byte))
		// fmt.Println(byte)
		print(byte)
		fmt.Sprintf("%s", p)
	}

	fmt.Println("\n" + string(p))

	return &s
}
