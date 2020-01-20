package hr

//Candidate 候选人
type Candidate interface {
	GetName() string
	GetCareer() string
	Accept(HR)
}

//Impl 候选人实体类
type Impl struct {
	Name   string
	Career string
}

//GetName 读取候选人姓名
//Return
//    string 候选人姓名
func (i *Impl) GetName() string {
	return i.Name
}

//GetCareer 读取候选人职业
//Return
//    string 候选人职业
func (i *Impl) GetCareer() string {
	return i.Career
}

//Accept accept
//Params
//    h HR hr
func (i *Impl) Accept(h HR) {
	h.Visit(i)
}
