<!--
 * @Date: 2022-01-26 14:08:16
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-26 14:52:27
 * @FilePath: \hrt014pocky\readme.md
-->
## 数组
```go
var 数组变量名 [元素数量]T
var a [5]int
```

## 切片
```go
var 切片变量名 []T
var name []T
```
### 使用make()函数构造切片
```go
make([]T, size, cap)
```
其中：
- T:切片的元素类型
- size:切片中元素的数量
- cap:切片的容量
