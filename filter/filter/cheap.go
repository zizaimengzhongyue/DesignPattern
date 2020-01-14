package filter

import "../candidate"

//Cheap 便宜候选人过滤器
type Cheap struct{}

//Filter 过滤便宜候选人
//Params
//    p []candidate.Candidate 候选人列表
//Return
//    []candidate.Candidate 符合条件的候选人列表
func (c Cheap) Filter(p []candidate.Candidate) []candidate.Candidate {
	res := []candidate.Candidate{}
	for _, v := range p {
		if v.GetSalary() <= 5000 {
			res = append(res, v)
		}
	}
	return res
}
