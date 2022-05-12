package myscanner

import (
	"io"
	"unicode/utf8"
)

type Scanner struct {
	// rune    rune
	text    string
	length  uint // 受け取った文字列長を格納
	current uint // 現在読み込んでいる文字位置を格納
	error   error
}

// スキャンする文字がまだある場合はtrueを返す
func (s *Scanner) Scan() bool {
	//文字を最後まで読み込んでいた場合はfalseを返す
	if s.current == s.length {
		return false
	}
	// エラーを保持している場合はfalseを返す
	if s.error != nil {
		return false
	}
	return true
}

// 現在の位置の文字を一つ返す
func (s *Scanner) Text() string {
	//s.current文字目（先頭が0）の文字を取り出す
	char := string(getRuneAt(s.text, s.current))
	s.current++
	return char
}

// io.Readerを受け取り、Scannerのポインタを返す
func NewScanner(r io.Reader) *Scanner {
	var s Scanner
	// スライスを確保
	var pl int = 100
	println(pl)
	tmp := make([]byte, pl)
	var p []byte
	var len int = pl
	var err error

	for pl == len {
		len, err = r.Read(tmp)

		if err != nil {
			// エラーが起こった場合はstructのerrorに値を保持し処理を終了させる。
			s.error = err
			return nil
		}

		// 取得した値をpに追加する
		p = append(p, tmp[:len]...)

	}
	// strに直してscannerに代入
	txt := string(p)
	s.text = txt
	// テキストの文字数を代入
	s.length = uint(utf8.RuneCountInString(txt))
	// fmt.Println("\n" + string(p))

	return &s
}

func (s Scanner) Err() error {
	return s.error
}

func getRuneAt(s string, i uint) rune {
	rs := []rune(s)
	return rs[i]
}
