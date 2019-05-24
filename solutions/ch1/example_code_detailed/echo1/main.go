// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// 源码详解者：chen sheng   date：2019年5月25日
// 只用于学习和交流，不做为商业用途


// 在terminal命令行运行main.go文件时，输出相应的参数信息
package main

import (
	"fmt"
	"os"  //go语言底层为我们解决了不依赖操作系统的操作API，即调用操作系统内核函数时，不需要区别是Linux或者是Windows了
)

//以下程序，使用terminal命令行操作效果比较明显
// 输入：go  run  路径/main.go
func main() {
	var s, sep string  //默认空值是""

	//Args是string类型的切片，如：[]string
	//切片就是slice，是array数组的一个view试图，可以对视图(slice)进行操作[1:],[:],[:5]
	//slice是一个前闭后开的区间，写是[1:5]，其实是[1:5) = [1:4]
	//args1是绝对路径下的文件名，即程序名，如：包名.exe
	for i := 0; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
