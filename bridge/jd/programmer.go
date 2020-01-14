package jd

import "fmt"

//Programmer 程序员
type Programmer struct{}

//JobDescription 工作描述
//Params
//    age int 年龄
//    salary int 薪资
//    experience 工作经验
func (p Programmer) JobDescription(age int, salary int, experience int) {
	description := "我们希望招聘一名程序员同学，要求年龄 %d，有 %d 年工作经验，月薪为 %d\n"
	fmt.Printf(description, age, experience, salary)
}
