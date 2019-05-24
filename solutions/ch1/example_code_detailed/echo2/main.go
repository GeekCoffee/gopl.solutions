// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// 源码详解者：chen sheng   date：2019年5月25日
// 只用于学习和交流，不做为商业用途

// 在terminal命令行运行main.go文件时，输出相应的参数信息
// 使用_ 和rang配合
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	//这个程序Windows命令行下是没有结果输出的，我稍微修改了一点
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)


	//使用strings.Join(os.Args[:], " ")合并更加高效
	//[:]表示从[0:len]所有的元素
	for _, arg := range os.Args[:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

