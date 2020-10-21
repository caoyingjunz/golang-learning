package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Url = "http://www.baidssu.com"
)

func main() {
	client := &http.Client{
		// TODO: 可以在此处设置超时或者其他
	}
	// 发送一个 POST 请求
	req, err := http.NewRequest("POST", Url, strings.NewReader("key=value"))
	if err != nil {
		fmt.Printf("NewRequest error: %v", err)
		return
	}

	// 增加header 或者 cookes (可选)
	req.Header.Add("Content-Type", "application/json")

	// 增加 cookes (可选)
	cookie1 := &http.Cookie{Name: "name", Value: "caoyingjun", HttpOnly: true}
	req.AddCookie(cookie1)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// 必须加 close 去关闭 连接
	// 需要在 err 处理后面，有的时候错误返回的时候，resp 为 nil
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Print(string(data))
}
