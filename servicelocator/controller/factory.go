package controller

//Factory 控制器工厂类
type Factory struct{}

//GetController 获取控制器
//Params
//    name string 名称
func (f *Factory) GetController(name string) Controller {
	if name == "controller01" {
		return &Controller01{}
	}
	return &Controller02{}
}
