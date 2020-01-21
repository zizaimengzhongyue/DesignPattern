package business

//LookUp 查找服务
type LookUp struct{}

//GetService 获取服务
//Params
//    name string 服务姓名
//Return
//    Service 服务
func (l *LookUp) GetService(name string) Service {
	switch name {
	case "it":
		return &IT{}
	case "hr":
		return &HR{}
	default:
		return nil
	}
}
