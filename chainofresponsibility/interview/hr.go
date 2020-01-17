package interview

import (
	"fmt"

	"../resume"
)

//HR 负责沟通面试安排的 HR
type HR struct {
	Next Interview
}

//SetNext 设置责任链中下一个状态的处理单位
func (h *HR) SetNext(i Interview) {
	h.Next = i
}

//Deal 处理面试
func (h *HR) Deal(r resume.Resume) {
	if r.GetStatus() == resume.STATUS_COMMUCATION {
		fmt.Printf("你好，%s，我是公司的 hr，负责与您沟通后续的面试事宜\n", r.GetName())
	} else {
		h.Next.Deal(r)
	}
}
