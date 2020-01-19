package memento

//Memento 备份单元数据
type Memento struct {
	content string
}

//SetContent 设置 content
//@params
//    content string  写入文本
func (m *Memento) SetContent(content string) {
	m.content = content
}

//GetContent 读取文本
//@Return
//    string 读取的文本
func (m *Memento) GetContent() string {
	return m.content
}
