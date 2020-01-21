package dao

//List 订单列表接口
type List interface {
	AddOrder(Order)
	UpdateOrder(Order)
	GetAllOrder(Order)
}

//Impl 订单列表实体类
type Impl struct {
	List []Order
}

//AddOrder 增加订单
//Params
//    o Order 订单
func (i *Impl) AddOrder(o Order) {
	i.List = append(i.List, o)
}

//UpdateOrder 更新订单信息，如果订单不存在则增加
//Params
//    o Order 订单
func (i *Impl) UpdateOrder(o Order) {
	for k, v := range i.List {
		if v.GetName() == o.GetName() {
			i.List[k] = o
			return
		}
	}
	i.AddOrder(o)
}

//GetAllOrder 展示所有订单信息
func (i *Impl) GetAllOrder() {
	for _, v := range i.List {
		v.ShowOrder()
	}
}
