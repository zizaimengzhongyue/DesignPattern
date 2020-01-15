package interview

import "fmt"

//Programmer 程序员面试
type Programmer struct {
	BaseInterview
}

//SecondInterview 第二轮面试
func (p Programmer) SecondInterview() {
	fmt.Println("这是第二轮面试")
}
