package employee

import "fmt"

//Employee 员工信息
type Employee struct {
	Name         string
	Age          int
	Subordinates []*Employee
}

//Description 介绍自己以及自己全部下属
func (e *Employee) Description() {
	fmt.Printf("我的名字是 %s，我今年 %d 岁\n", e.Name, e.Age)
	for _, v := range e.Subordinates {
		v.Description()
	}
}

//AddSubordidate 增加下属
//Params
//    m *Employee 员工信息
func (e *Employee) AddSubordidate(m *Employee) {
	e.Subordinates = append(e.Subordinates, m)
}
