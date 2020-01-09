//Description
//  对于一场面试，不同的面试者面试成功或者失败之后会发送不同的消息，比如对于程序员则是“欢迎加入技术部”，销售则是“欢迎加入销售部”，将 interview 的 Pass 和 Reject 具体执行通过其 candidate 属性的 Pass 和
//Reject 方法来执行，这样可以针对不同的面试者发送不同的消息
//  如果后续需要加入行政、人力等不同的岗位，只要创建相应的行政类和人力类并分别实现 Pass 和 Reject 方法即可实现发送新的消息
package main

import (
	"./candidate"
)

type Candidate interface {
	Pass()
	Reject()
}

type interview struct {
	candidate Candidate
}

func (i interview) Pass() {
	i.candidate.Pass()
}

func (i interview) Reject() {
	i.candidate.Reject()
}

func main() {
	programmer := candidate.Programmer{}
	programmerInterview := interview{candidate: programmer}
	programmerInterview.Pass()
	programmerInterview.Reject()
	salesman := candidate.Salesman{}
	salesmanInterview := interview{candidate: salesman}
	salesmanInterview.Pass()
	salesmanInterview.Reject()
}
