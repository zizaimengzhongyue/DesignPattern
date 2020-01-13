package interview

import (
	"./candidate"
	"./interviewer"
)

//Builder 创建者
type Builder struct{}

//Programmer 创建程序员面试
//Return
//    Interview 面试
func (b Builder) Programmer() Interview {
	programmerInterviewer := interviewer.Programmer{}
	programmerCandidater := candidate.Programmer{}
	return Interv{interv: programmerInterviewer, canda: programmerCandidater}
}

//Salesman 创建销售面试
//Return
//    Interview 面试
func (b Builder) Salesman() Interview {
	salesmanInterviewer := interviewer.Salesman{}
	salesmanCandidate := candidate.Salesman{}
	return Interv{interv: salesmanInterviewer, canda: salesmanCandidate}
}
