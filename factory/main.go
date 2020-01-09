//Description
//  我们希望根据需要邀请不同的候选人来面试。
//  这里可以通过一个工厂类实现，根据输入的职业名称创建不同的候选人
package main

import "./candidate"

func main() {
	factory := candidate.Factory{}

	programmer := factory.GetCandidate("programmer")
	programmer.AboutMyself()

	salesman := factory.GetCandidate("salesman")
	salesman.AboutMyself()
}
