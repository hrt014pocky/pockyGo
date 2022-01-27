/*
 * @Date: 2022-01-27 23:50:24
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 23:57:35
 * @FilePath: \hrt014pocky\package_demo\dsa\dsa.go
 */
package main

// 当代码都放在GOPATH目录下面
// 导入包要从&GOPATH/src后面继续写起
import (
	"fmt"

	"github.com/hrt014pocky/package_demo/calc"
	// "github.com/hrt014pocky/package_demo/calc"
)

func main() {
	fmt.Println("hello")
	n := calc.Name
	fmt.Println("Person name is", n)
}
