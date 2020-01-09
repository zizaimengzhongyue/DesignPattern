package candidates

import "fmt"

//Programmer 程序员
type Programmer struct{}

//AboutMyself 自我介绍
func (p Programmer) AboutMyself() {
	fmt.Println("我是一名程序员")
}
