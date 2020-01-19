package interview

import "fmt"

//Reject 面试失败
type Reject struct {
	Interv *Interviewer
}

//Execute 面试失败
func (r *Reject) Execute() {
	fmt.Printf("%s 先生/女士，很遗憾您没有通过我们的面试。\n", r.Interv.GetName())
}
