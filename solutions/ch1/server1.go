// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// 版权所有者是Alan A. A. Donovan & Brian W. Kernighan.


package main

import (
	"fmt"
	"log"
	"net/http"
)

// 做一个迷你web服务器，功能是：发送一个get或者post请求给Web服务器，返回给客户端请求的路径信息
func main(){

	//给请求处理程序进行注册,把handler处理函数和"/"开头的URL绑定在一起
	//表示以“/”开头的URL都需要用handler函数来处理，如: /hello, /hello/golang , /*
	http.HandleFunc("/",handler)

	//把源代码改写了一下
	//开启服务器监听程序，监听"ip:8000"端口下是否有请求，返回一个err对象
	//第一个参数是要监听的全路径地址，如：192.168.10.10:8080
	//第二个参数不是下面的自定义函数func，是一个Hander结构体的对象
	err := http.ListenAndServe("localhost:8000",nil)
	if err != nil{
		log.Fatal(err)
	}

}

// request处理程序
// 参数是不需要我们传入的，是go语言接收到client的请求后，把数据包封装在http.Request对象里 ，帮我们传入的
// 注意*http.Request是加了一个*的，说明是引用传递，是一个真正的从客户端传来的request对象
func handler(w http.ResponseWriter, r *http.Request){

	//w是一个返回体，需要做什么事情就写到w里面返回给客户端
	//用fmt.Fprintf方法把数据写到返回体w里，当函数handler执行完后，w会被返回给client端
	fmt.Fprintf(w,"URL.path = %q\n",r.URL.Path)
}
