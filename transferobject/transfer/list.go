package transfer

import "fmt"

//List 名单
type List struct {
	Employees []Employee
}

//Add 增加员工
//Params
//    e Employee 员工
func (l *List) Add(e Employee) {
	l.Employees = append(l.Employees, e)
}

//Fire 解雇一名员工
//Params
//    e Employee 员工信息
func (l *List) Fire(e Employee) {
	index := -1
	for k, v := range l.Employees {
		if v.GetName() == e.GetName() {
			index = k
			break
		}
	}
	if index != -1 {
		l.Employees = append(l.Employees[0:index], l.Employees[index+1:]...)
	}
}

//Show 展示员工信息
func (l *List) Show() {
	for _, v := range l.Employees {
		fmt.Printf("姓名：%s，薪资：%d\n", v.GetName(), v.GetSalary())
	}
}
