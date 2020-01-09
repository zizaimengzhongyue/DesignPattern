package interviewers

import "fmt"

//Programmer 程序员面试官
type Programmer struct{}

//AboutMyself 自我介绍
func (p Programmer) AboutMyself() {
	fmt.Println("我是一名程序员面试官")
}
