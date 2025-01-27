package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` 提供原始命令行参数访问功能。注意，切片中
	// 的第一个参数是该程序的路径，并且 `os.Args[1:]`保存
	// 所有程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
