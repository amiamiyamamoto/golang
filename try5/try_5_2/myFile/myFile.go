package myFile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type File string

// .jpgもしくは.jpegファイルだった場合はtrueを返す
func (f File) IsJpg() bool {
	e := f.GetExt()
	if e == ".jpg" || f == ".jpeg" {
		return true
	}
	return false
}

// func (f *File) changeJpgToPng()

// 拡張子を返す関数をつくる
// IsJpgでこの関数を使う
// changeJpgToPngメソッドをつくる→実際のファイル名操作、自身のファイル名を変更する
func (f File) GetExt() string {
	e := filepath.Ext(string(f))
	return e
}

func PrintFile(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			PrintFile(dir + "/" + file.Name())
			continue
		}
		var cfile File = File(dir + "/" + file.Name())
		fmt.Println(cfile.IsJpg())

		fmt.Println(dir + "/" + file.Name())

	}
}
