package composite

//Object 接口
type Object interface {
	SetData(string)
	GetData()
}

//Impl object 实体类
type Impl struct {
	Data string
}

//SetData 设置 data
//Params
//    data string 写入内容
func (i *Impl) SetData(data string) {
	i.Data = data
}

//GetData
//Return
//    string 读取内容
func (i *Impl) GetData() string {
	return i.Data
}
