package expression

//Expression 解释器 interface
type Expression interface {
	Interpret(interface{}) bool
}
