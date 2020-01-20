package candidate

import "fmt"

//Salesman  销售
type Salesman struct {
	Impl
}

//AboutMe 自我介绍
func (s *Salesman) AboutMe() {
	fmt.Printf("我的名字是 %s，我的职业是销售。\n", s.Name)
}
