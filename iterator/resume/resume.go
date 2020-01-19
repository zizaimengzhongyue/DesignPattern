package resume

//Resume 简历
type Resume interface {
	GetName() string
}

//Impl Resume 实体类
type Impl struct {
	Name string
}

//GetName 获取简历姓名
//Return
//    string 简历姓名
func (r *Impl) GetName() string {
	return r.Name
}
