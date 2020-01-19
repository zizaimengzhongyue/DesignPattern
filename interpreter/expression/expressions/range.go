package expressions

//Range 匹配区间的解释器
type Range struct {
	Min int
	Max int
}

//Interpret 判断是否符合条件
//Params
//    context 上下文
//Return
//    bool 是否在 [rangeMin, rangeMax] 这个区间内
func (r *Range) Interpret(context interface{}) bool {
	vContext, ok := context.(int)
	if !ok {
		return false
	}
	return vContext >= r.Min && vContext < r.Max
}
