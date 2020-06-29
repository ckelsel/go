package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// 建一个通道ch，用于跟main线程通讯
	ch := make(chan string)

	// Get所有数据
	for _, url := range os.Args[1:] {
		// 启动协程
		go fetch(url, ch)
	}

	// 打印所有数据
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	} else {
		// TODO
		ch <- fmt.Sprintf("Get %s success", url)
        resp.Body.Close()
	}
}
