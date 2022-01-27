/*
 * @Date: 2022-01-27 16:25:32
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 17:01:38
 * @FilePath: \hrt014pocky\interfaces\p_v.go
 */

package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 使用值类型实现接收者接口：类型的值和类型的指针都能够保存在接口的变量中
func (p Person) move() {
	fmt.Printf("%v is moving\n", p.name)
}

func (p Person) say() {
	fmt.Printf("%v is saying\n", p.name)
}
func (p Person) shout() {
	fmt.Printf("%v is shouting\n", p.name)
}
func (p Person) run() {
	fmt.Printf("%v is running\n", p.name)
}
func (p Person) jog() {
	fmt.Printf("%v is jogging\n", p.name)
}

// 使用指针类型实现接收者接口：只有类型指针能够保存到变量接口
// func (p *Person) move() {
// 	fmt.Printf("%vis moving\n", p.name)
// }

type mover interface {
	move()
	run()
	jog()
}
type sayer interface {
	say()
	shout()
}

//接口的嵌套
type allow interface {
	mover
	sayer
}

func main() {
	// p1 := Person{
	// 	name: "Alice",
	// 	age:  20,
	// }
	p2 := &Person{
		name: "Bob",
		age:  22,
	}
	var m mover
	var s sayer
	var a allow
	// m = p1
	// m.move()
	m = p2
	m.move()
	m.run()
	m.jog()
	s = p2
	s.say()
	s.shout()
	a = p2
	a.say()

}
