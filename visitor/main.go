//Description
//    对于同一个候选人的信息可能有多个部门的 HR 来看
package main

import "./hr"

func main() {
	object := hr.Object{}
	candidate01 := &hr.Impl{Name: "张三", Career: "程序员"}
	object.Add(candidate01)
	candidate02 := &hr.Impl{Name: "李四", Career: "销售"}
	object.Add(candidate02)

	visitorSales := hr.Sales{}
	object.Accept(visitorSales)
	visitorTechnology := hr.Technology{}
	object.Accept(visitorTechnology)
}
