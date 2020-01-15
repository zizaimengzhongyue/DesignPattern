package interview

import "fmt"

//Interview 面试抽象类
type Interview interface {
	FirstInterview()
	LastInterview()
}

//BaseInterview 面试基类
type BaseInterview struct{}

//FirstInterview 第一轮面试
func (bi BaseInterview) FirstInterview() {
	fmt.Println("这是第一轮面试")
}

//LastInterview 最后一轮面试
func (bi BaseInterview) LastInterview() {
	fmt.Println("这是最后一轮面试")
}
