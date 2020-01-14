package jd

import "fmt"

//Salesman 销售
type Salesman struct{}

//JobDescription 工作描述
//Params
//    age int 年龄
//    salary int 薪资
//    experience int 工作经验
func (s Salesman) JobDescription(age int, salary int, experience int) {
	description := "我们希望招聘一名销售同学，要求年龄 %d，有 %d 年工作经验，月薪为 %d\n"
	fmt.Printf(description, age, experience, salary)
}
