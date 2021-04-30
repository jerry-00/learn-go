package main

import (
	"fmt"
)

//变量使用注意事项
func main() {
	//该区域的数值可以在同一类型范围内不断变化
	var i = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	//在初始化变量声明变量时候就给值
	var a int = 30
	//在声明时候赋值 省略数据类型
	var a1 int = 30
	fmt.Println(a, a1)
	//先声明 后赋值
	var num int
	fmt.Println(num)
}
