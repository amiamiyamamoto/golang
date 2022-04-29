package main

type Stringer interface {
	String() string
}

func main() {

}

// 任意の値をStringer型に変換する関数
func ToStringer(v interface{}) (Stringer, error) {

}
