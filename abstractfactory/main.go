//Description
//  我们要安排一场面试，我们有针对不同职位的不同面试官和候选人，我们希望根据职位不同分别获取面试官和候选人
//  在写抽象工厂模式的时候遇到一些问题
//  1. 对于面试官工厂应该直能获得面试官，从面试官工厂不能获取候选人，候选人工厂同理。但是创建的工厂的 interface 定义了获取面试官和候选人两个接口，
//这就需要在两个工厂中分别覆写一边一个用不到的方法，并且返回 nil。如果后续增加新的职位类型，那就需要在工厂的 interface 里面新增加一个接口并且修改所有已有的工厂类；
//  2. 不同的工厂因为要重复实现获取其他角色的方法，这又需要在每个工厂类中引用其他角色的定义。这种跨上级目录引用不合理而且容易出现循环引用的问题。
//  第一个问题，我定义了一个工厂类的基类，在这个基类里面实现了获取不同职位的方法并且返回 nil，其他工厂继承该基类，工厂只要覆写自己能创建的角色的方法就行。
//  第二个问题，我把所有的工厂类放在同一个目录下面来解决这个问题
//  对抽象工厂模式我应该还有什么误会，看起来好像只是个工厂的工厂。多了解之后再做修改。
package main

import "./interview"

func main() {
	abstractFactory := interview.AbstractFactory{}

	interviewerFactory := abstractFactory.GetFactory("interviewer")
	programmerInterviewer := interviewerFactory.GetInterviewer("programmer")
	programmerInterviewer.AboutMyself()
	salesmanInterviewer := interviewerFactory.GetInterviewer("salesman")
	salesmanInterviewer.AboutMyself()

	candidateFactory := abstractFactory.GetFactory("candidate")
	programmerCandidate := candidateFactory.GetCandidate("programmer")
	programmerCandidate.AboutMyself()
	salesmanCandidate := candidateFactory.GetCandidate("salesman")
	salesmanCandidate.AboutMyself()
}
