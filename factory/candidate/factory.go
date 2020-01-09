package candidate

import "./candidates"

//Candidate 候选人
type Candidate interface {
	AboutMyself()
}

//Factory 候选人工厂类
type Factory struct{}

func (c Factory) GetCandidate(name string) Candidate {
	switch name {
	case "programmer":
		return candidates.Programmer{}
	case "salesman":
		return candidates.Salesman{}
	default:
		return nil
	}
}
