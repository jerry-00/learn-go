//表示改hello.go文件所在的包是main 在go中 每个文件都必须归属一个包
package main

//引入一个包 包名fmt 引入后可使用fmt包函数
import "fmt"

//func 是个关键字 表示后面是一个函数
//编译 build hello.go 生成可执行文件
func main() {
	fmt.Println("hello,world!")
}
