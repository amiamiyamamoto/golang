package myscanner

import (
	"fmt"
	"io"
)

type Scanner struct {
	// rune    rune
	data    []byte
	reader  io.Reader //いらないかも
	length  uint      // 受け取った文字列長を格納
	current uint      // 現在読み込んでいる文字位置を格納
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
func NewScanner(r io.Reader) (*Scanner, error){
	var s Scanner
	s.reader = r
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
			return nil err
		}

		// 取得した値をpに追加する
		p = append(p, tmp[:len]...)

		fmt.Printf("%+v\n %+v\n\n\n\n\n", p, err)
		// fmt.Println(tmp)

	}
	//TODO:取得した[]byteを　string → rune に変換してScannerに代入する
	fmt.Println(len)

	for _, byte := range p {
		// fmt.Println(string(byte))
		// fmt.Println(byte)
		print(byte)
		fmt.Sprintf("%s", p)
	}

	fmt.Println("\n" + string(p))

	return &s
}
