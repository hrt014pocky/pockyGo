/*
 * @Date: 2022-01-26 12:36:23
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 12:43:14
 * @FilePath: \hrt014pocky\if\main.go
 */

package main

import "fmt"

func main() {
	// 1.基本写法
	// var score = 65
	// if score >= 90 {
	//   fmt.Println("A")
	// } else if score >= 75 {
	//   fmt.Println("B")
	// } else {
	//   fmt.Println("C")
	// }
	// 2. if的特殊写法, score只在if语句块内生效
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
