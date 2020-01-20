package hr

import "fmt"

//Sales 销售部 hr
type Sales struct{}

//Visit visit 接口
//Params
//    i *Imp 候选人
func (s Sales) Visit(i *Impl) {
	if i.GetCareer() == "销售" {
		fmt.Printf("%s 一名销售，这就是我们想要的\n", i.GetName())
	} else {
		fmt.Printf("%s 不是一名销售，这不是我们想要的\n", i.GetName())
	}
}
