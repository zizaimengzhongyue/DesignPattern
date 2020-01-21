package controller

//FilterManager 过滤管理器
type FilterManager struct {
	FilterChain
}

//AddFilter 增加过滤器
func (f *FilterManager) AddFilter(fi Filter) {
	f.FilterChain.AddFilter(fi)
}

//Execute 处理请求
//Params
//    request string 请求
func (f *FilterManager) Execute(request string) {
	f.FilterChain.Execute(request)
}
