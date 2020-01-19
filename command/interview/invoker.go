package interview

//Interviewer 面试者
type Interviewer struct {
	Name string
}

//GetName 获取面试者的姓名
//Return
//    string 面试者姓名
func (i *Interviewer) GetName() string {
	return i.Name
}
