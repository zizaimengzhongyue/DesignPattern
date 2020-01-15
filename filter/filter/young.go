package filter

import "../candidate"

//Young 年轻候选人过滤器
type Young struct{}

//Filter 过滤年轻人
//Params
//    p []candidate.Candidate 年轻候选人名单
//Return
//    []candidate.Candidate 符合条件候选人名单
func (y Young) Filter(p []candidate.Candidate) []candidate.Candidate {
	res := []candidate.Candidate{}
	for _, v := range p {
		if v.GetAge() <= 35 {
			res = append(res, v)
		}
	}
	return res
}
