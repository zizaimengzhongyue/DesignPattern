package controller

import "fmt"

//Controller01 控制器
type Controller01 struct{}

//GetName 获取名称
//Return
//    string 名称
func (c *Controller01) GetName() string {
	return "controller01"
}

//Handle 处理请求
//Params
//    request string 请求
func (c *Controller01) Handle(request string) {
	fmt.Printf("Controller01 handle %s\n", request)
}
