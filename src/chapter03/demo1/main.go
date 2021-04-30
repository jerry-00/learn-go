package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义变量
	var i int
	//赋值
	i = 10
	//使用
	fmt.Println("i=", i)
	fmt.Printf("n1的类型是", i)
	var n2 int64 = 10
	//可以返回n2变量占用的字节数
	fmt.Println()
	fmt.Printf("n2占用的字节", n2, unsafe.Sizeof(n2))
	//go 使用过程中遵守保小不保大的原则
	//保证程序在正确运行下,尽量占用空间小的数据类型

}
