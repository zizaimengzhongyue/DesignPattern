package interview

import "fmt"

//Second 第二轮面试
type Second struct{}

//Deal 处理
//Params
//    c *Candidate 候选人
func (s *Second) Deal(c *Candidate) {
	fmt.Printf("%s 先生/女士，您第二轮面试的结果是 %s\n", c.GetName(), c.GetResult())
}
