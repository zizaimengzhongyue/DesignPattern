package department

import (
	"../candidate"
	"../talents"
)

//Technology 技术部
type Technology struct {
	Talents *talents.Talents
}

//ShowCandidate 展示候选人信息
//    *candidate.Impl 候选人
func (t *Technology) ShowCandidate(c *candidate.Impl) {
	t.Talents.ShowCandidate(c)
}
