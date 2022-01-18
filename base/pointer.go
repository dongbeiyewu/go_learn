package main

import "fmt"

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	num := 100
	add(num)
	fmt.Println(num) // 100，num 没有变化

	realAdd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改
}
