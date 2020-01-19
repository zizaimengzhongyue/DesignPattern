package food

//Real 存在的食物
type Real struct {
	Name string
}

//GetName 食物名称
//Return
//    string 食物名称
func (r *Real) GetName() string {
	return r.Name
}

//IsNil 是否为 nil
//Return
//    bool 是否为 nil true 为是，false 为否
func (r *Real) IsNil() bool {
	return false
}
