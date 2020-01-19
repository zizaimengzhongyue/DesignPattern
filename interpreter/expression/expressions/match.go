package expressions

//Match 字符串匹配解释器
type Match struct {
	Str string
}

//Interpret 校验
//Params
//    context interface{} 上下文
//Return
//    bool 是否匹配规则
func (m *Match) Interpret(context interface{}) bool {
	vContext, ok := context.(string)
	if !ok {
		return false
	}
	return vContext == m.Str
}
