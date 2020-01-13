package interview

//Interviewer 面试官
type Interviewer interface {
	Question()
}

//Candidater 候选人
type Candidater interface {
	Answer()
}

//Interview 面试抽象类
type Interview interface {
	Question()
	Answer()
}

//Interv 一场面试包括两个角色，面试官和候选人
type Interv struct {
	interv Interviewer
	canda  Candidater
}

//Question 面试问题
func (i Interv) Question() {
	i.interv.Question()
}

//Answer 面试答案
func (i Interv) Answer() {
	i.canda.Answer()
}
