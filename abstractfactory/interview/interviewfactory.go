package interview

import "./interviewers"

//Interviewer 面试官
type Interviewer interface {
	AboutMyself()
}

//InterviewerFactory 面试官工厂
type interviewerFactory struct {
	BaseFactory
}

//GetInterviewer 创建面试官
func (inf interviewerFactory) GetInterviewer(name string) Interviewer {
	switch name {
	case "programmer":
		return interviewers.Programmer{}
	case "salesman":
		return interviewers.Salesman{}
	default:
		return nil
	}
}
