//Description
//    很多路由分发就是这么处理的
package main

import "./controller"

func main() {
	client := controller.Client{}
	hello := &controller.Hello{}
	client.AddFilter(hello)
	world := &controller.World{}
	client.AddFilter(world)
	client.Execute("hello")
}
