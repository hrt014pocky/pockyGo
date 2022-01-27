/*
 * @Date: 2022-01-27 15:15:53
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 15:43:44
 * @FilePath: \hrt014pocky\struct\main.go
 */
package main

import (
	"fmt"
)

type Person struct {
	name, city string
	age        int
}

func main() {
	p1 := new(Person)
	p1.name = "Alice"
	p1.city = "beijing"
	p1.age = 18
	fmt.Printf("p1=%#v\n", p1)

}
