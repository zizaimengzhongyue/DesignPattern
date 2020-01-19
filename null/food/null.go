package food

//Null 不存在的食物
type Null struct{}

//GetName 食物名称
//Return
//    string 食物名称
func (n *Null) GetName() string {
	return "菜品不存在"
}

//IsNil 是否为 nil
//Return
//    bool 是否为 nil true 为是，false 为否
func (n *Null) IsNil() bool {
	return true
}
