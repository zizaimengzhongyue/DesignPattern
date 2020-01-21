package controller

import "fmt"

//Controller02 控制器
type Controller02 struct{}

//Handle 处理请求
func (c *Controller02) Handle() {
	fmt.Println("Controller02 handle")
}
