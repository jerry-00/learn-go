//表示改hello.go文件所在的包是main 在go中 每个文件都必须归属一个包
package main

//引入一个包 包名fmt 引入后可使用fmt包函数
import "fmt"

//func 是个关键字 表示后面是一个函数
//编译 build hello.go 生成可执行文件
//在Windows下编译生成指定运行文件 go build -o xxx.exe xxx.go linux 可直接go run 编译运行
//与java不同每一行后面不需要加分号 go是一行一行编译 不要把多条语句放一起写
/*
注意事项 go官方推荐使用行注释
*/
func main() {
	fmt.Println("hello,world!")
	//go import包没有使用到编译会报错 变量声明未使用也会编译报错
	//var name = 10
}
