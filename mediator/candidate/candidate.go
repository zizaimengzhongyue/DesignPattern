package candidate

//Candidate 候选人接口
type Candidate interface {
	GetName() string
	GetCarrer() string
}

//Impl 候选人基类
type Impl struct {
	Name   string
	Carrer string
}

//GetName 获取候选人姓名
//Return
//    string 候选人姓名
func (i *Impl) GetName() string {
	return i.Name
}

//GetCarrer 获取候选人职业
//    候选人职业
func (i *Impl) GetCarrer() string {
	return i.Carrer
}
