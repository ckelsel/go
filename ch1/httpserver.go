package main

import (
"fmt"
"net/http")

func main() {
    // 1. 绑定/的处理函数
    http.HandleFunc("/", handler)

    // 2. 运行服务
    http.ListenAndServe("localhost:8000", nil)
}

// 3. 实现处理函数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
