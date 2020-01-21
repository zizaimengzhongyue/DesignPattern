package controller

import "fmt"

//Controller 控制器
type Controller struct{}

//Handle 处理请求
func (c *Controller) Handle() {
	fmt.Println("Controller handle")
}
