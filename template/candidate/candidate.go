package candidate

import "fmt"

//Candidate 候选人
type Candidate interface {
	GetName() string
	GetCarrer() string
	AboutMe()
}

//Impl 候选人实体类
type Impl struct {
	Name string
}

//GetName 读取候选人姓名
//Return
//    string 候选人姓名
func (i *Impl) GetName() string {
	return i.Name
}

//AboutMe 自我介绍
func (i *Impl) AboutMe() {
	fmt.Printf("我的名字是 %s。\n", i.Name)
}
