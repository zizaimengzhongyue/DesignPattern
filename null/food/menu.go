package food

//Menu 菜单
type Menu struct {
	List []Food
}

//Add 菜单增加新菜品
//Params
//    f Food 菜品
func (m *Menu) Add(f Food) {
	m.List = append(m.List, f)
}

//GetFood 读取菜品
//Params
//    name string 菜品名
//Return
//    Food 菜品
func (m *Menu) GetFood(name string) Food {
	for _, v := range m.List {
		if v.GetName() == name {
			return v
		}
	}
	return &Null{}
}
