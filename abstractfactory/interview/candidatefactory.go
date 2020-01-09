package interview

import "./candidates"

//Candidate 候选人
type Candidate interface {
	AboutMyself()
}

//Factory 候选人工厂类
type candidateFactory struct {
	BaseFactory
}

//GetCandidate 获取候选人
func (cf candidateFactory) GetCandidate(name string) Candidate {
	switch name {
	case "programmer":
		return candidates.Programmer{}
	case "salesman":
		return candidates.Salesman{}
	default:
		return nil
	}
}
