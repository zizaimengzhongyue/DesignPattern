//Description
//   看起来只是将一个对象传来传去，对象只能有简单的 set/get 方法。这样就算序列化之后也不会丢失信息
package main

import "./transfer"

func main() {
	list := transfer.List{}
	employee01 := transfer.Employee{Name: "张三", Salary: 5000}
	list.Add(employee01)
	employee02 := transfer.Employee{Name: "李四", Salary: 6000}
	list.Add(employee02)
	list.Show()
}
