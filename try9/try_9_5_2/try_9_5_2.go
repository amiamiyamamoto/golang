package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	url := "https://1.bp.blogspot.com/-tVeC6En4e_E/X96mhDTzJNI/AAAAAAABdBo/jlD_jvZvMuk3qUcNjA_XORrA4w3lhPkdQCNcBGAsYHQ/s1048/onepiece01_luffy.png"
	fn := "sample.png"
	if err := download(url, fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func download(url string, fn string) error {

	// rangeで分割して複数回ダウンロードする
	var r uint = 500000000 // 分割するサイズ
	var c uint = 0         // 現在位置
	// var resp *http.Response

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		// リクエストを飛ばしてデータを取得
		resp, err := rangeReq(url, r, c)
		c = c + r + 1
		if err != nil {
			return err
		}
		if resp.StatusCode != int(206) {
			break
		}

		// 取得したデータをファイルに書き込む
		if err := postscript(fn, resp); err != nil {
			return err
		}

		// f.Write([]byte(a))
		// f.Write(resp.Body)
	}

	// _, err = io.Copy(f, resp.Body)
	return nil
}

func postscript(fn string, resp *http.Response) error {
	a := resp.Body

	file, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	// file.Write(a)
	fmt.Fprintln(file, a) //書き込みってこれであってんの？？なにこれ
	return nil

}

func rangeReq(url string, r uint, c uint) (*http.Response, error) {
	client := &http.Client{ // CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Range", "bytes="+strconv.FormatUint(uint64(c), 10)+"-"+strconv.FormatUint(uint64(c+r), 10))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	return resp, nil
}
