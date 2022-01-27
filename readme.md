<!--
 * @Date: 2022-01-26 14:08:16
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 15:00:15
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
## map
Go语言中 map的定义语法如下：
```go
map[KeyType]ValueType
```
其中，

KeyType:表示键的类型。
ValueType:表示键对应的值的类型。
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
```go
make(map[KeyType]ValueType, [cap])
```

## new和make
1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。