/*
 * @Date: 2022-01-24 16:32:31
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 12:09:25
 * @FilePath: \hrt014pocky\studyGo\main.go
 */
package main

import "fmt"

const pi = 3.1415

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "Alice"
	age = 18
	isOk = true

	fmt.Printf("she's name is %s", name)
	fmt.Println()
	fmt.Print(name)
	fmt.Println("'s age is", age)
	fmt.Println("is ok?", isOk)
	fmt.Println("pi is ", pi)
}
