/*
 * @Date: 2022-01-26 15:19:05
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 15:40:03
 * @FilePath: \hrt014pocky\map\main.go
 */
package main

import "fmt"

func main() {
	// map的初始化
	scoreMap := make(map[string]int, 8)
	// map中如何添加键值对
	scoreMap["Alice"] = 80
	scoreMap["Bob"] = 50
	fmt.Println(scoreMap)
	// 声明map的同时完成初始化
	passMap := map[string]bool{
		"Alice": true,
		"Bol":   false,
	}
	fmt.Println(passMap)
	fmt.Printf("type:%T\n", passMap)
	// 判断某个键值是否存在 value, ok := map[key]
	fmt.Println("==========================value ok")
	v, ok := scoreMap["Carol"]
	if ok {
		fmt.Println("Socre is", v)
	} else {
		fmt.Println("Score map doesnt have this person!")
	}
}
