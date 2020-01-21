package controller

//FilterChain filter chain
type FilterChain struct {
	List []Filter
	Controller
}

//AddFilter 增加过滤器
//Params
//   fi Filter 过滤器
func (f *FilterChain) AddFilter(fi Filter) {
	f.List = append(f.List, fi)
}

//Execute 执行过滤
//Params
//    request string 请求
func (f *FilterChain) Execute(request string) {
	for _, v := range f.List {
		v.Execute(request)
	}
	f.Controller.Handle()
}
