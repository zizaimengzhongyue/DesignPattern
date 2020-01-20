package hr

import "fmt"

//Technology 技术部 hr
type Technology struct{}

//Visit visit 接口
//Params
//    i *Imp 候选人
func (t Technology) Visit(i *Impl) {
	if i.GetCareer() == "程序员" {
		fmt.Printf("%s 是一名程序员，这就是我们想要的\n", i.GetName())
	} else {
		fmt.Printf("%s 不是一名程序员，这不是我们想要的\n", i.GetName())
	}
}
