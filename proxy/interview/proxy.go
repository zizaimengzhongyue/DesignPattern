package interview

//Interview 面试接口
type Interview interface {
	JobDescription()
}

//Proxy Interview 的代理类
type Proxy struct {
	Interv Interview
}

//JobDescription 面试职位描述
func (p *Proxy) JobDescription() {
	p.Interv.JobDescription()
}
