//Description
//    我们定义不同的面试官对于面试者进行面试，如果通过就执行 Pass 方法，如果失败就执行 Reject 方法。
//    命令模式、策略模式、状态模式分不清楚
package main

import "./interview"

func main() {
	interviewer01 := &interview.Interviewer{Name: "张三"}
	executer01 := interview.Pass{Interv: interviewer01}
	executer01.Execute()

	interviewer02 := &interview.Interviewer{Name: "李四"}
	executer02 := interview.Reject{Interv: interviewer02}
	executer02.Execute()
}
