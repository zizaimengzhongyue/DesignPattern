package interview

//Interview 面试接口
type Interview interface {
	Deal(Candidate)
}

//Candidate 候选人
type Candidate struct {
	Name   string
	Result string
}

//GetName 读取候选人姓名
//Return
//    string 候选人姓名
func (c *Candidate) GetName() string {
	return c.Name
}

//GetResult 读取候选人面试结果
//Return
//    string 候选人面试结果
func (c *Candidate) GetResult() string {
	return c.Result
}
