package main

import "fmt"

func main() {

	// 要创建一个空 map，需要使用内建的 `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// 使用典型的 `make[key] = val` 语法来设置键值对。
	m["k1"] = 7
	m["k2"] = 13

	// 使用例如 `fmt.Println` 来打印一个 map 将会输出所有的
	// 键值对。
	fmt.Println("map:", m)

	// 使用 `name[key]` 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 当对一个 map 调用内建的 `len` 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	// 内建的 `delete` 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键
	// 是否在这个 map 中。这可以用来消除键不存在和键有零值，
	// 像 `0` 或者 `""` 而产生的歧义。这里我们不需要值，所以
	// 用_空白标识符(blank identifier)_忽略。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 你也可以通过这个语法在同一行声明和初始化一个新的
	// map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
