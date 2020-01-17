//Description
//    一份简历存在待沟通，面试中，面试通过，薪资沟通等一系列连续的状态，不同的状态需要不同的人来处理。通过责任链模式将这个状态在责任链中传递处理
package main

import (
	"./interview"
	"./resume"
)

func buildChainOfResponsibility() interview.Interview {
	hr := &interview.HR{}
	programmer := &interview.Programmer{}
	hrd := &interview.HRD{}
	hr.SetNext(programmer)
	programmer.SetNext(hrd)
	return hr
}

func main() {
	interv := buildChainOfResponsibility()
	resume01 := &resume.Programmer{Name: "张三", Status: resume.STATUS_PASS}
	interv.Deal(resume01)
	resume02 := &resume.Programmer{Name: "李四", Status: resume.STATUS_INTERVIEW}
	interv.Deal(resume02)
}
