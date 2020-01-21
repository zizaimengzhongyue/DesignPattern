package controller

import "fmt"

//Controller01 控制器
type Controller01 struct{}

//Handle 处理请求
func (c *Controller01) Handle() {
	fmt.Println("Controller01 handle")
}
