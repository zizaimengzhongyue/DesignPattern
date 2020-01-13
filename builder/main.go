//Description
//  创建一场面试，我们需要安排面试官和候选人，以面试为纬度，针对程序员面试和销售面试构造不同的 builder 类
//  我还没搞明白建造者模式，也没想到有什么工作中遇到过的场景适用这种设计模式，实现可能也有问题。后续再修改补充
package main

import (
	"./interview"
)

func main() {
	builder := interview.Builder{}

	programmerInterview := builder.Programmer()
	programmerInterview.Question()
	programmerInterview.Answer()

	salesmanInterview := builder.Salesman()
	salesmanInterview.Question()
	salesmanInterview.Answer()
}
