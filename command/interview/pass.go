package interview

import "fmt"

//Pass 面试通过执行的命令类
type Pass struct {
	Interv *Interviewer
}

//Execute 面试通过
func (p *Pass) Execute() {
	fmt.Printf("%s 先生/女士，恭喜您通过我们的面试。\n", p.Interv.GetName())
}
