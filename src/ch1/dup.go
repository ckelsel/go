package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFromStdin() {
	// 定义一个map, string:int
	counts := make(map[string]int)

	// 从标准输入读取数据
	input := bufio.NewScanner(os.Stdin)

	// 如果数据重复，则map对应的值加1
	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1
	}

	fmt.Println("-------------")
	// 输入结束后，输出重复的数据，>1
	for key, value := range counts {
		if value > 1 {
			fmt.Println(key)
			fmt.Println(value)
		}
	}
}

func main() {
	counts := make(map[string]int)

	if len(os.Args) <= 1 {
		return
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		return
	}

	input := bufio.NewScanner(f)

	// 如果数据重复，则map对应的值加1
	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1
	}

	fmt.Println("-------------")
	// 输入结束后，输出重复的数据，>1
	for key, value := range counts {
		if value > 1 {
			fmt.Println(key)
			fmt.Println(value)
		}
	}

	f.Close()
}
