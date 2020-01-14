//Description
//    本身是个树形结构，定义的 struct 跟实现一棵树的 struct 也没什么区别
package main

import (
	"./employee"
)

func main() {
	employee01 := employee.Employee{Name: "张三", Age: 30}
	employee02 := employee.Employee{Name: "张四", Age: 29}
	employee03 := employee.Employee{Name: "张五", Age: 28}
	employee04 := employee.Employee{Name: "张六", Age: 27}
	employee05 := employee.Employee{Name: "张七", Age: 26}
	employee06 := employee.Employee{Name: "张八", Age: 25}
	employee07 := employee.Employee{Name: "张九", Age: 24}

	employee01.AddSubordidate(&employee02)
	employee01.AddSubordidate(&employee03)
	employee02.AddSubordidate(&employee04)
	employee02.AddSubordidate(&employee05)
	employee03.AddSubordidate(&employee06)
	employee03.AddSubordidate(&employee07)

	employee01.Description()
}
