package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	client := &http.Client{}
	// 发送一个请求
	req, err := http.NewRequest("POST", "http://163.com", strings.NewReader("key=value"))
	defer req.Body.Close()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	// 必须加 close
	req.Header.Add("User-Agent", "myClient")

	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(data))
}
