//Description
//    通过 mvc 描述一名雇员的信息
package main

import "./employee"

func main() {
	view := &employee.View{}
	emp := &employee.Employee{Name: "张三", Salary: 5000}
	controller := &employee.Controller{View: view, Emp: emp}

	controller.Show()
	controller.UpdateEmployeeSalary(10000)
	controller.Show()
}
