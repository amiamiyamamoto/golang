package main

import (
	"fmt"
	"io"
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
	var r uint = 100000 // 分割するサイズ
	var c uint = 0      // 現在位置
	// var resp *http.Response

	// status code 406でforを終了させる
	for _, _ = range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		resp, err := rangeReq(url, r, c)
		if err != nil {
			return err
		}
		if resp.StatusCode == 406 {
			break
		}
	}

	out, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func rangeReq(url string, r uint, c uint) (*http.Response, error) {
	client := &http.Client{ // CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Range", "bytes="+strconv.Itoa(c)+"-"+strconv.Itoa(c+r))
	c = c + r + 1

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	return resp, nil
}
