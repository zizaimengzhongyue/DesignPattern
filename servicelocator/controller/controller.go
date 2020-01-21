package controller

//Controller Controller 接口
type Controller interface {
	GetName() string
	Handle(string)
}
