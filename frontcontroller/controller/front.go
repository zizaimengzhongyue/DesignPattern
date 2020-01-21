package controller

import "fmt"

//Front 前端控制器
type Front struct {
	Dispatch
}

//Intercept 拦截器
func (f *Front) Intercept() {
	fmt.Println("Front intercept")
}

//DispatchRequest 分发请求
//Params
//    name string 请求
func (f *Front) DispatchRequest(name string) {
	f.Intercept()
	f.Dispatch.Dispatch(name)
}
