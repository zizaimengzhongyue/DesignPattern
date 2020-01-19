package department

import (
	"../candidate"
	"../talents"
)

//Sales 销售部
type Sales struct {
	Talents *talents.Talents
}

//ShowCandidate 展示候选人信息
//    *candidate.Candidate 候选人
func (s *Sales) ShowCandidate(c *candidate.Impl) {
	s.Talents.ShowCandidate(c)
}
