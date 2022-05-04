package myscanner

import "io"

type Scanner struct {
	text   string
	reader io.Reader
}

func NewScanner(r io.Reader) *Scanner {
	// inf := ^uint(0)
	// for i := uint(666); i < inf; i++ {
	// 	println("tako")
	// }
	var s Scanner
	s.reader = r
	return &s
}
