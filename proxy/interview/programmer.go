package interview

import "fmt"

//Programmer 程序员面试
type Programmer struct{}

//JobDescription 工作描述
func (p Programmer) JobDescription() {
	fmt.Println("这是一份 GoLang 程序员的岗位，要求应聘者需要有 GoLang 相关开发经验，熟悉 Go 语言")
}
