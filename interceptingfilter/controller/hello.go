package controller

import "fmt"

//Hello fitler
type Hello struct{}

//Execute 过滤
//Params
//    request string 请求
func (h *Hello) Execute(request string) {
	if request == "hello" {
		fmt.Println("这个请求是 hello")
	} else {
		fmt.Println("这个请求不是 hello")
	}
}
