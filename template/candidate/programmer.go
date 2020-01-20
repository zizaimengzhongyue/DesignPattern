package candidate

import "fmt"

//Programmer 程序员
type Programmer struct {
	Impl
}

//AboutMe 自我介绍
func (p *Programmer) AboutMe() {
	fmt.Printf("我的名字是 %s，我的职业是程序员。\n", p.Name)
}
