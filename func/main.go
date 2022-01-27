/*
 * @Date: 2022-01-26 20:17:34
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 20:55:06
 * @FilePath: \hrt014pocky\func\main.go
 */
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello!")
}

func add(x int, y int) int {
	return x + y
}

// 函数接收可变参数
// 可变参数在函数体重是一个切片
func add1(x ...int) int {
	res := 0
	for _, arg := range x {
		res += arg
	}
	return res
}

// 匿名函数和闭包
// 定义一个函数,它的返回值是一个函数
func hello1() func() {
	name := "pocky"
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	sayHello()
	sum := add(1, 3)
	fmt.Println(sum)
	sum = add1(1, 2, 3, 4)
	fmt.Println(sum)
	// 闭包 = 函数 + 外层变量的引用
	f := hello1()
	f()
}
