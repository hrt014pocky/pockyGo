/*
 * @Date: 2022-01-26 12:43:44
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 13:11:13
 * @FilePath: \hrt014pocky\for\main.go
 */
package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// for i < 15 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("hello pocky!")
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		break
	// 	}
	// }
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
