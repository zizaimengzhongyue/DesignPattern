package interview

//Factory 工厂
type Factory interface {
	GetCandidate(name string) Candidate
	GetInterviewer(name string) Interviewer
}

//BaseFactory 工厂基类
type BaseFactory struct{}

//GetCandidate 获取候选人
func (bf BaseFactory) GetCandidate(name string) Candidate {
	return nil
}

//GetInterviewer 获取面试者
func (bf BaseFactory) GetInterviewer(name string) Interviewer {
	return nil
}

//AbstractFactory 抽象工厂
type AbstractFactory struct{}

//GetFactory 获取工厂
func (af AbstractFactory) GetFactory(name string) Factory {
	switch name {
	case "candidate":
		return candidateFactory{BaseFactory{}}
	case "interviewer":
		return interviewerFactory{BaseFactory{}}
	default:
		return nil
	}
}
