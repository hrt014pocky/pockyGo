/*
 * @Date: 2022-01-27 14:24:37
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 14:28:50
 * @FilePath: \hrt014pocky\pointer\main.go
 */
package main

import "fmt"

func main() {
	a := 10 // int类型
	b := &a // int*类型
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, &b)
}
