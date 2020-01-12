package sington

var interview interviewID

type interviewID struct {
	ID int
}

func (i *interviewID) GetID() int {
	i.ID++
	return i.ID
}

//GetInstance 获取实例
//Return
//    *interviewID 单例的实例
func GetInstance() *interviewID {
	return &interview
}

func init() {
	interview = interviewID{}
}
