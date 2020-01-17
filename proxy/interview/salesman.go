package interview

import "fmt"

//Salesman 销售职位
type Salesman struct{}

//JobDescription 工作描述
func (s *Salesman) JobDescription() {
	fmt.Println("这是一份销售的工作，需要应聘者语言表达能力强，有销售相关的工作经验")
}
