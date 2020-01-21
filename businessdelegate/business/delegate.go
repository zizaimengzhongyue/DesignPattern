package business

//DeleGate 代表
type DeleGate struct {
	LookUp
	ServiceType string
}

//SetServiceType 设置服务类型
//Params
//    name string 服务类型
func (d *DeleGate) SetServiceType(name string) {
	d.ServiceType = name
}

//DoTask 执行任务
func (d *DeleGate) DoTask() {
	d.LookUp.GetService(d.ServiceType).DoProcessing()
}
