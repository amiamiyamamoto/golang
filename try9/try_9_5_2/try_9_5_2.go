package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {

	url := "http://www.kms.ac.jp/~clinilab/cgi-bin/makeweb/sample.txt"
	fn := "sample.txt"
	if err := download(url, fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func download(url string, fn string) error {

	// リクエストを飛ばしてデータを取得
	var i int64 = 0
	for {
		if i > 900 {
			break
		}
		go tmp(i, url, fn)

		// resp, err := rangeReq(url, uint(i), uint(i+300))
		// i = i + 1 + 300
		// if err != nil {
		// 	return err
		// }
		// //データを取得し終えたらforを離脱
		// if resp.StatusCode != int(206) {
		// 	break
		// }
		// // ファイルに書き込み
		// if err := postscript(fn, resp); err != nil {
		// 	return err
		// }
		i = i + 1 + 300
	}

	return nil
}

var wg = sync.WaitGroup

func tmp(i int64, url string, fn string) {
	resp, err := rangeReq(url, uint64(i), uint64(i+300))
	i = i + 1 + 300
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	//データを取得し終えたらforを離脱
	if resp.StatusCode != int(206) {
		// break
	}
	// ファイルに書き込み
	if err := postscript(fn, resp); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func postscript(fn string, resp *http.Response) error {

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	// もしファイルがshift-jisでもutf-8に変換可能
	// utf8_body, err := encodeUtf8(body)
	// if err != nil {
	// 	return err
	// }

	file, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(body)
	return nil

}

// もしファイルがshift-jisでもutf-8に変換可能
func encodeUtf8(byte []byte) ([]byte, error) {
	utf8_body := transform.NewReader(bytes.NewReader(byte), japanese.ShiftJIS.NewDecoder())
	return ioutil.ReadAll(utf8_body)
}

func rangeReq(url string, s uint64, e uint64) (*http.Response, error) {
	client := &http.Client{ // CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Range", "bytes="+strconv.FormatUint(s, 10)+"-"+strconv.FormatUint(e, 10))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
