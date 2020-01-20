package hr

// Object visitor 控制 visitor 访问 元素
type Object struct {
	List []*Impl
}

//Accept accept
//Params
//    h HR hr
func (o *Object) Accept(h HR) {
	for _, v := range o.List {
		v.Accept(h)
	}
}

//Add 增加候选人
//c *Impl 候选人
func (o *Object) Add(c *Impl) {
	o.List = append(o.List, c)
}
