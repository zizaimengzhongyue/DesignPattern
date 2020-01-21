package dao

import "fmt"

//Order 订单
type Order struct {
	Name  string
	Price int
}

//SetName 设置名称
//Params
//    name string 订单名称
func (o *Order) SetName(name string) {
	o.Name = name
}

//GetName 读取订单名称
//Return
//    string 订单名称
func (o *Order) GetName() string {
	return o.Name
}

//SetPrice 设置订单价格
//Params
//    price int 订单价格
func (o *Order) SetPrice(price int) {
	o.Price = price
}

//GetPrice 读取订单价格
//Return
//    price int 订单价格
func (o *Order) GetPrice(price int) int {
	return o.Price
}

//ShowOrder 展示订单信息
func (o *Order) ShowOrder() {
	fmt.Printf("订单名称：%s，订单价格：%d\n", o.Name, o.Price)
}
