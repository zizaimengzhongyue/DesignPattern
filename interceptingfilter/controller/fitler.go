package controller

//Filter 过滤器
type Filter interface {
	Execute(string)
}
