package controller

//Controller 控制器接口
type Controller interface {
	Handle()
}

//Dispatch 路由分发器
type Dispatch struct {
	Controller01
	Controller02
}

//Dispatch 分发请求
//Params
//    request string 请求
func (d *Dispatch) Dispatch(request string) {
	if request == "controller01" {
		d.Controller01.Handle()
	} else {
		d.Controller02.Handle()
	}
}
