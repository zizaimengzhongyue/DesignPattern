package candidates

import "fmt"

//Salesman 销售
type Salesman struct{}

//AboutMyself 自我介绍
func (s Salesman) AboutMyself() {
	fmt.Println("我是一名销售")
}
