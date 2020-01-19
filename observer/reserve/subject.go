package reserve

//Subject 观察的项目
type Subject struct {
	observers []Observer
	price     int
}

//Attach 绑定观察者
//Params
//    o Observer 观察者
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

//Update 更新价格
//Params
//    price int 价格
func (s *Subject) Update(price int) {
	s.price = price
	s.NotifyObservers()
}

//GetPrice 读取价格
//Return
//    int 价格
func (s *Subject) GetPrice() int {
	return s.price
}

//NotifyObservers 通知观察者
func (s *Subject) NotifyObservers() {
	for _, v := range s.observers {
		v.Update()
	}
}
