//Description
//    大部分服务的路由分发都是这么搞的
package main

import "./controller"

func main() {
	front := &controller.Front{}
	front.DispatchRequest("controller01")
	front.DispatchRequest("controller02")
}
