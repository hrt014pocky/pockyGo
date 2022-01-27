/*
 * @Date: 2022-01-27 17:03:32
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 17:27:10
 * @FilePath: \hrt014pocky\interfaces\null_interfance\main.go
 */
package main

import "fmt"

//空接口
// 接口中没有定义任何需要实现的方法时, 该接口就是一个空接口

// 空接口的应用
// 1. 空接口类型作为函数的参数, 函数可以接收任意类型的变量
// 2. 空接口可以作为map的value

func main() {
	var x interface{}
	x = "hello"
	x = 100

	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "Alice"
	m["age"] = 18
	m["hobby"] = []string{"唱", "跳", "rap", "篮球"}
	fmt.Println(m)

	ret, ok := x.(string)
	if ok {
		fmt.Println(ret)
	} else {
		fmt.Println("This is not string type")
	}

}
