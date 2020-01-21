//Description
//    我看 demo 里是在 list 的构造方法里面初始化学生信息，我这里是客户端手动初始化
package main

import "./dao"

func main() {
	list := &dao.Impl{}
	order1 := dao.Order{Name: "手机", Price: 1999}
	list.AddOrder(order1)
	order2 := dao.Order{Name: "电脑", Price: 7999}
	list.AddOrder(order2)
	list.GetAllOrder()

	order2.SetPrice(6999)
	list.UpdateOrder(order2)
	list.GetAllOrder()
}
