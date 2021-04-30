package main

import "fmt"

//定义全局变量
//var n1 = 101
//var n2 = 105
//var name = "jack"
//改成一次性声明
var (
	n5    = 500
	n6    = 600
	name2 = "mary"
)

func main() {
	//go 一次声明多个变量
	//var n1, n2, n3 int
	//fmt.Println(n1, n2, n3)

	//一次声明多个变量方式 类型不同
	//var n1, name ,n3 = 100, "tom", 99
	//fmt.Println("n1=",n1,"name=",name,"n3=",n3)
	//n1= 100 name= tom n3= 99

	//一次声明多种变量 类型推导 相当于用 := 分割吧
	n1, name, n3 := 100, "tom", 99
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)
	//n1= 100 name= tom n3= 99
	fmt.Println("n5=", n5, "name2=", name2, "n6=", n6)
}
