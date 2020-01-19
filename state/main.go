//Description
//    一个候选人参加面试，根据候选人当前面试在第几轮输出不同的结果
//    状态模式和策略模式很像
package main

import "./interview"

func main() {
	candidate := &interview.Candidate{Name: "张三", Result: "通过"}
	first := &interview.First{}
	second := &interview.Second{}
	last := &interview.Last{}
	first.Deal(candidate)
	second.Deal(candidate)
	last.Deal(candidate)
}
