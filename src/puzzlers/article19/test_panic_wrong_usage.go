package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")
}

func caller() {
	fmt.Println("Enter function caller.")
	panic(errors.New("something wrong")) // 正例。
	// 这个recover函数调用并不会起到任何作用，甚至都没有机会执行
	p := recover() //
	fmt.Printf("panic: %s\n", p)
}
