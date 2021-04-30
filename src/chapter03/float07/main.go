package main

import "fmt"

//go 中小数类型使用
func main() {

	var cc float32 = -39.2
	fmt.Println(cc)
	//精度更高一点
	var ccc float64 = 39.333
	fmt.Println(ccc)
	//go默认使用float64
	var ddddd = .3333333333333333
	fmt.Println(ddddd)
	//科学计数法 e2 10^2
	var ccccc = .3333333e2
	fmt.Println(ccccc)
	// 除 10^2
	var dddddd = .33333e-2
	fmt.Println(dddddd)
}
