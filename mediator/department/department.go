package department

import "../candidate"

//Department 部门接口
type Department interface {
	ShowCandidate(*candidate.Impl)
}
