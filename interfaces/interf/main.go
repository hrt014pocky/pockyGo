/*
 * @Date: 2022-01-27 15:44:06
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 16:51:09
 * @FilePath: \hrt014pocky\interfaces\interf\main.go
 */
package main

import "fmt"

type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type sayer interface {
	say()
}

func beat(arg sayer) {
	arg.say()
}

func main() {
	var c1 cat
	var d1 dog
	beat(c1)
	beat(d1)
}
