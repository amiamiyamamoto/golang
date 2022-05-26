package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "https://1.bp.blogspot.com/-tVeC6En4e_E/X96mhDTzJNI/AAAAAAABdBo/jlD_jvZvMuk3qUcNjA_XORrA4w3lhPkdQCNcBGAsYHQ/s1048/onepiece01_luffy.png"
	fn := "sample.png"
	if err := download(url, fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func download(url string, fn string) error {
	client := &http.Client{ // CheckRedirect: redirectPolicyFunc,
	}

	// rangeで分割して複数回ダウンロードする
	// 0-500, 501-1000, 1001-1500
	for _, num := range []int{1, 2, 3} {

	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Range", "bytes=9999999990-9999999999")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
