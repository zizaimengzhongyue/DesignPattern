package controller

import "fmt"

//World filter
type World struct{}

//Execute 执行过滤
//Params
//    request string 请求
func (w *World) Execute(request string) {
	if request == "world" {
		fmt.Println("这个请求是 world")
	} else {
		fmt.Println("这个请求不是 world")
	}
}
