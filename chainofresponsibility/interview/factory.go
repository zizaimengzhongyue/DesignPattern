package interview

import "../resume"

//Interview 面试接口 interface
type Interview interface {
	SetNext(Interview)
	Deal(resume.Resume)
}
