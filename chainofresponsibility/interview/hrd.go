package interview

import (
	"fmt"

	"../resume"
)

//HRD 负责沟通薪资和员工入职
type HRD struct {
	Next Interview
}

//SetNext 设置责任链中下一个状态的处理单位
func (h *HRD) SetNext(i Interview) {
	h.Next = i
}

//Deal 处理面试
func (h *HRD) Deal(r resume.Resume) {
	if r.GetStatus() == resume.STATUS_PASS {
		fmt.Printf("你好，%s，我是公司的 hrd，负责与您沟通薪资和入职的事宜\n", r.GetName())
	} else {
		h.Next.Deal(r)
	}
}
