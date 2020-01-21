package controller

//Locator 定位器
type Locator struct {
	Factory
	Cache
}

//GetController 获取控制器
//Params
//    name string 名称
//Return
//    *Controller 控制器
func (l *Locator) GetController(name string) Controller {
	v := l.Cache.GetController(name)
	if v != nil {
		return v
	}

	con := l.Factory.GetController(name)
	l.Cache.SetController(con)
	return con
}
