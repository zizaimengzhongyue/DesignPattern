package memento

//Originator struct
type Originator struct {
	content string
}

//SetContent 写入文本
//Params
//    content string 写入的文本
func (o *Originator) SetContent(content string) {
	o.content = content
}

//GetContent 读出文本
//Return
//    string 读出文本
func (o *Originator) GetContent() string {
	return o.content
}

//SaveContentToMemento 保存文本到 memento
//Return
//    *Memento 保存文本的 Memento
func (o *Originator) SaveContentToMemento() *Memento {
	return &Memento{content: o.content}
}

//RestoreContentFromMemento 从 Memento 恢复数据
//Params
//    m *Memento 备份的 Memento
func (o *Originator) RestoreContentFromMemento(m *Memento) {
	o.content = m.GetContent()
}
