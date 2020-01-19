package memento

//CareTaker 备份数据
type CareTaker struct {
	list []*Memento
}

//Add 备份
//Params
//    m *Memento 备份的数据
func (c *CareTaker) Add(m *Memento) {
	c.list = append(c.list, m)
}

//Get 读取备份
//Params
//    index int 备份数据下标
//Return
//    *Memento 备份数据
func (c *CareTaker) Get(index int) *Memento {
	if index < len(c.list) {
		return c.list[index]
	}
	return nil
}
