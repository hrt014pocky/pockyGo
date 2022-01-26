/*
 * @Date: 2022-01-26 13:26:48
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 13:53:01
 * @FilePath: \hrt014pocky\array\main.go
 */
package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	// 编译器推到数组的长度
	var cityAmerica = [...]string{"NewYork", "LosAngle", "Hoston", "SanFrancisco"}
	fmt.Println(cityAmerica)
	fmt.Println(cityAmerica[3])

	var nbaTeam = [...]string{1: "Lakers", 3: "Nets", 6: "Warriors", 7: "Rockets"}
	fmt.Println(nbaTeam)
	fmt.Printf("%T\n", nbaTeam)

	// 数组的遍历
	// 1.使用for循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	// 2.for range 遍历
	for index, value := range cityArray {
		fmt.Println(index, value)
	}
	for _, value := range cityArray {
		fmt.Println(value)
	}

	cities := [2][3]string{
		{"北京", "上海", "广州"},
		{"NewYork", "LosAngle", "Hoston"},
	}
	fmt.Println(cities)
	fmt.Println(cities[1][0])
	// 二维数组的遍历
	for _, v1 := range cities {
		fmt.Println(v1)
	}
	for _, v1 := range cities {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
