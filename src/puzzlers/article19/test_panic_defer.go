package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	// 当panic引发时调用defer函数
	// 尽量把defer语句写在函数体的开始处，因为在引发 panic 的语句之后的所有语句，都不会有任何执行机会
	defer func() {
		fmt.Println("Enter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())

	// 引发panic。
	panic(errors.New("something wrong"))

	// panic后面的语句不会被执行
	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}
