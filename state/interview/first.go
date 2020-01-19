package interview

import "fmt"

//First 第一轮面试
type First struct{}

//Deal 处理
//Params
//    c *Candidate 候选人
func (f *First) Deal(c *Candidate) {
	fmt.Printf("%s 先生/女士，您第一轮面试的结果是 %s\n", c.GetName(), c.GetResult())
}
