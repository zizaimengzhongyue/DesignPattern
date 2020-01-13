package resume

import "fmt"

//Resume 简历
type Resume struct {
	name   string
	mobile string
}

//SetName  设置简历姓名
//Params
//    name string 姓名
func (r *Resume) SetName(name string) {
	r.name = name
}

//GetName 输出简历姓名
func (r *Resume) GetName() {
	fmt.Println(r.name)
}

//SetMobile 设置简历电话
//Params
//    mobile string 简历电话
func (r *Resume) SetMobile(mobile string) {
	r.mobile = mobile
}

//GetMobile 输出简历电话
func (r *Resume) GetMobile() {
	fmt.Println(r.mobile)
}

//Clone 复制简历
//Return
//    *Resume 返回简历的 copy
func (r *Resume) Clone() *Resume {
	temp := *r
	return &temp
}
