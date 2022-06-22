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
	resp, err := rangeReq(url, 0, 50000000)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	return nil
}

func postscript(fn string, resp *http.Response) error {
	a := resp.Body
	defer a.Close()

	file, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	// file.Write(a)
	byte, err := ioutil.ReadAll(a)
	if err != nil {
		return err
	}

	fmt.Println(byte)
	fmt.Println(a)
	file.Write(byte)
	// fmt.Fprintln(file, a)
	fmt.Fprintln(file, "aaaaa")
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
