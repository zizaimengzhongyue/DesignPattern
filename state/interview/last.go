package interview

import "fmt"

//Last 最终面
type Last struct{}

//Deal 处理
//Params
//    c *Candidate 候选人
func (s *Last) Deal(c *Candidate) {
	fmt.Printf("%s 先生/女士，您最终面试的结果是 %s\n", c.GetName(), c.GetResult())
}
