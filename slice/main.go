/*
 * @Date: 2022-01-26 14:04:26
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 15:12:39
 * @FilePath: \hrt014pocky\slice\main.go
 */
package main

import "fmt"

func main() {
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	fmt.Println(a, b, c)
	fmt.Println(a == nil)
	// 基于数组得到切片
	numsArray := [5]int{1, 2, 3, 4, 5}
	numsSlice := numsArray[1:4]
	fmt.Println(numsSlice)
	// 从切片再获得切片
	numsSlice1 := numsSlice[0:]
	fmt.Println(numsSlice1)
	// make构造切片
	fmt.Println("==========================make")
	sm := make([]int, 5, 10)
	fmt.Println(sm)
	fmt.Printf("%T\n", sm)
	fmt.Println(len(sm))
	fmt.Println(cap(sm))
	//nil
	fmt.Println("==========================nil")
	var king []int      //声明int类型切片
	var queen = []int{} //声明并且初始化
	var jack = make([]int, 0)
	fmt.Println(king, len(king), cap(king))
	fmt.Println("king is nil:", (king == nil))
	fmt.Println(queen, len(queen), cap(queen))
	fmt.Println("queen is nil:", (queen == nil))
	fmt.Println(jack, len(jack), cap(jack))
	fmt.Println("jack is nil:", (jack == nil))
	// 要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	fmt.Println("==========================copy")
	//切片的赋值拷贝
	//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	coffee := []string{"Americano", "Latte", "Mocha"}
	cafe := coffee
	cafe[0] = "Cappuccino"
	fmt.Println(coffee)
	// 切片的扩容
	fmt.Println("==========================append")
	// 切片要初始化之后才能使用
	var ap []int // 此时并没有申请内存
	ap = append(ap, 10)
	fmt.Println(ap)
	ap[0] = 100
	fmt.Println(ap)
	for i := 0; i < 5; i++ {
		ap = append(ap, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", ap, len(ap), cap(ap), ap)
	}
	ap = append(ap, 1, 2, 3, 4, 5)
	fmt.Println(ap)
	ap1 := []int{11, 12, 13}
	ap = append(ap, ap1...)
	fmt.Println(ap)

}
