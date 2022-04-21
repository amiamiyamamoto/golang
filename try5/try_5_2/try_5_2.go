package main

import (
	"flag"
	"fmt"
	myfile "myfile/myFile"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	// 指定したディレクトリ以下のJPGファイルをPNGに変換するコマンドを作成する
	// 自作パッケージ
	// ユーザ定義型を作成
	// GoDocを生成
	// Go Modulesを使う
	myfile.PrintFile(flag.Args()[0])

}
