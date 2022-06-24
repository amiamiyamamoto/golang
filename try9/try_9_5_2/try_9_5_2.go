package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	i := 0
	for {
		resp, err := rangeReq(url, uint(i), uint(i+300))
		i = i + 1 + 300
		if err != nil {
			return err
		}
		//データを取得し終えたらforを離脱
		if resp.StatusCode != int(206) {
			break
		}
		// ファイルに書き込み
		if err := postscript(fn, resp); err != nil {
			return err
		}
	}

	return nil
}

func postscript(fn string, resp *http.Response) error {

	body, err := ioutil.ReadAll(resp.Body)
	// _ = body
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	file, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println(resp.Header["Content-Type"])
	file.Write(body)
	// fmt.Fprintln(file, body)
	// io.Copy(file, resp.Body)
	return nil

}

func rangeReq(url string, s uint, e uint) (*http.Response, error) {
	client := &http.Client{ // CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Range", "bytes="+strconv.FormatUint(uint64(s), 10)+"-"+strconv.FormatUint(uint64(e), 10))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
