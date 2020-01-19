package expression

//AndExpression And 类型解释器
type AndExpression struct {
	Expressions []Expression
}

//Add 增加解释器
//Params
//    e Expression 解释器
func (ae *AndExpression) Add(e Expression) {
	ae.Expressions = append(ae.Expressions, e)
}

//Interpret 判断一组 context 是否都符合条件
//Params
//    contexts []interface{} 一组 context
//Return
//    bool 是否满足全部条件
func (ae *AndExpression) Interpret(contexts []interface{}) bool {
	ans := true
	i := 0
	for ; i < len(ae.Expressions) && i < len(contexts); i++ {
		ans = ans && ae.Expressions[i].Interpret(contexts[i])
	}
	return ans && (i == len(ae.Expressions))
}
