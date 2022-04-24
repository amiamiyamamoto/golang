package myFile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type File string

// .jpgもしくは.jpegファイルだった場合はtrueを返す
func (f File) isJpg() bool {
	e := f.getExt()
	if e == ".jpg" || e == ".jpeg" {
		return true
	}
	return false
}

// ファイルがjpgもしくはjpegだった場合にpngファイルに変換する
func (f *File) changeJpgToPng() error {
	// TODO:エラーの処理を追加する
	if !f.isJpg() {
		return nil
	}
	nfn := strings.Replace(string(*f), f.getExt(), ".png", 1)
	if err := os.Rename(string(*f), nfn); err != nil {
		return err
	}
	fmt.Println("convert: " + string(*f))
	*f = File(nfn)
	return nil

}

// file型の拡張子をピリオド付きで返す
func (f File) getExt() string {
	e := filepath.Ext(string(f))
	return e
}

// 指定したディレクトリのjpgもしくはjpegファイルをpngファイルに変換する
func ChangeFilesJpgToPng(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, file := range files {
		if file.IsDir() {
			ChangeFilesJpgToPng(dir + "/" + file.Name())
			continue
		}
		var cfile File = File(dir + "/" + file.Name())
		if err := cfile.changeJpgToPng(); err != nil {
			return err
		}
	}
	return nil
}
