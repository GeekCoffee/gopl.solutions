// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// 源码详解者：chen sheng   date：2019年5月25日
// 只用于学习和交流，不做为商业用途


// 在terminal命令行运行main.go文件时，输出相应的参数信息
// 使用strings.Join(os.Args[:], " ")合并更加高效
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//这个程序Windows命令行下是没有结果输出的，我稍微修改了一点
	//fmt.Println(strings.Join(os.Args[1:], " "))

	//这个可以输出绝对路径+文件名
	fmt.Println(strings.Join(os.Args[0:], " "))

	//以[...]形式输出slice
	fmt.Println(os.Args[:])
}


