package main

import (
	"net/http"
	"os"
	"sync"
)

var group sync.WaitGroup

func main() {

	// 建一个通道ch，用于跟main线程通讯
	// ch := make(chan string)

	// Get所有数据
	for _, url := range os.Args[1:] {

		// 1. 加入组
		group.Add(1)

		go fetch(url)

		// 启动协程
		// go fetch(url, ch)
	}

	// 3. 等待所有协程结束
	group.Wait()

	// 打印所有数据
	// for range os.Args[1:] {
	//     fmt.Println(<-ch)
	// }
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
	} else {
		resp.Body.Close()
	}

	group.Done()

	// 2. 从组中删除
}

// func fetch(url string, ch chan<- string) {
//     resp, err := http.Get(url)
//     if err != nil {
//         ch <- fmt.Sprint(err)
//     } else {
//         // TODO
//         ch <- fmt.Sprintf("Get %s success", url)
//         resp.Body.Close()
//     }
// }
