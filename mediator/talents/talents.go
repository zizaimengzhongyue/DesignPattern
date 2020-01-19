package talents

import (
	"fmt"

	"../candidate"
)

//Talents 人才库
type Talents struct{}

//ShowCandidate 展示候选人信息
//Params
//    *candidate.Candidate 候选人
func (t *Talents) ShowCandidate(c *candidate.Impl) {
	fmt.Printf("这名候选人的姓名是 %s，职业是 %s。\n", c.GetName(), c.GetCarrer())
}
