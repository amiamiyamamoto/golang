package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	url := "http://www.kms.ac.jp/~clinilab/cgi-bin/makeweb/sample.txt"
	num := []int{1, 2, 3, 4}
	data := make(map[int][]byte)
	for _, i := range num {
		wg.Add(1)
		go req(url, 0, 300, i, data)
	}
	wg.Wait()
	fmt.Println(data)

}

func req(url string, s uint64, e uint64, no int, data map[int][]byte) {
	defer wg.Done()

	// リクエストを投げる
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Range", "bytes="+strconv.FormatUint(s, 10)+"-"+strconv.FormatUint(e, 10))

	resp, _ := client.Do(req)
	// return resp

	//配列に追加する
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	data[no] = body

}
