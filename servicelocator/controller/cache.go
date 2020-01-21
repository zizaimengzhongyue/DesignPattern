package controller

//Cache 缓存
type Cache struct {
	List []Controller
}

//GetController 获取控制器
//Params
//    name string 控制器名称
//Return
//    Controller 控制器
func (c *Cache) GetController(name string) Controller {
	for _, v := range c.List {
		if v.GetName() == name {
			return v
		}
	}
	return nil
}

//SetController 设置控制器
//Params
//    name string 控制器名称
//    co Controller 控制器
func (c *Cache) SetController(co Controller) {
	c.List = append(c.List, co)
}
