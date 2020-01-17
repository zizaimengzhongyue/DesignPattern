package interview

import (
	"fmt"

	"../resume"
)

//Programmer  程序员面试官，负责面试
type Programmer struct {
	Next Interview
}

//SetNext 设置责任链中下一个状态的处理单位
func (h *Programmer) SetNext(i Interview) {
	h.Next = i
}

//Deal 处理面试
func (h *Programmer) Deal(r resume.Resume) {
	if r.GetStatus() == resume.STATUS_INTERVIEW {
		fmt.Printf("你好，%s，我是公司的面试官，负责您的面试工作\n", r.GetName())
	} else {
		h.Next.Deal(r)
	}
}
