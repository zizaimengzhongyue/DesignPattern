package controller

import "fmt"

//Controller02 控制器
type Controller02 struct{}

//GetName 获取名称
//Return
//    string 名称
func (c *Controller02) GetName() string {
	return "controller02"
}

//Handle 处理请求
//Params
//    request string 请求
func (c *Controller02) Handle(request string) {
	fmt.Printf("Controller02 handle %s\n", request)
}
