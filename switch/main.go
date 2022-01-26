/*
 * @Date: 2022-01-26 13:12:56
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 13:23:38
 * @FilePath: \hrt014pocky\switch\main.go
 */
package main

import "fmt"

func main() {
	num := 2
	switch num {
	case 1:
		fmt.Println("Alice")
	case 2:
		fmt.Println("Bob")
	case 3:
		fmt.Println("Carol")
	case 4:
		fmt.Println("Dave")
	default:
		fmt.Println("Invalid!")
	}

	switch num {
	case 1, 3:
		fmt.Println("Odd number")
	case 2, 4:
		fmt.Println("Even number")
	default:
		fmt.Println(num)
	}

	switch {
	case num <= 2:
		fmt.Println("Come on!")
	case 2 < num && num < 10:
		fmt.Println("Baby!")
	}
}
