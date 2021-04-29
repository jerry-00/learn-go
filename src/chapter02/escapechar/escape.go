package main

//主要提供格式化 输入输出函数
import "fmt"

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	//路径转义符
	fmt.Println("E:\\xunlei\\Mini Diva")
	fmt.Println("tom说\"i love you\"")
	//回车 \r后替换掉前面内容 新版本变成全覆盖 目前使用版本为1.16.3
	fmt.Println("xxxxxbbbbb\raaaa")
	fmt.Println("姓名\t性别\t年龄\nstan\t男\t18")
}
