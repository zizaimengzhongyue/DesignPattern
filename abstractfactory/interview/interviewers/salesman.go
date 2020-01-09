package interviewers

import "fmt"

//Salesman 销售面试官
type Salesman struct{}

//AboutMyself 自我介绍
func (s Salesman) AboutMyself() {
	fmt.Println("我是一名销售面试官")
}
