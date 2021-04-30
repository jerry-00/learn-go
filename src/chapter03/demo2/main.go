package main

import "fmt"

func main() {
	//定义变量
	var i int
	//赋值
	//i = 10
	//使用 int的默认值为0
	fmt.Println("i=", i)

	//第二种类型推导 根据值自行判定变量类型
	var num = 10.11
	fmt.Println("i=", num)

	//第三种省略var 注意:= 左侧的变量不应该是已经声明过的,否则会编译错误
	//下面的方式 相当于 var name int name = "20.5"   :不能省略
	name := 20.5
	fmt.Println("name=", name)
}
