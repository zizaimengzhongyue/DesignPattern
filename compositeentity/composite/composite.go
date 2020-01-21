package composite

//Composite 融合实体类
type Composite struct {
	Object1
	Object2
}

//SetData 写入数据
//Params
//    data1 string 写入实体 1 的数据
//    data2 string 写入实体 2 的数据
func (c *Composite) SetData(data1 string, data2 string) {
	c.Object1.SetData(data1)
	c.Object2.SetData(data2)
}

//GetData 读取数据
//Return
//    []string 读取的数据
func (c *Composite) GetData() []string {
	datas := []string{}
	datas = append(datas, c.Object1.GetData())
	datas = append(datas, c.Object2.GetData())
	return datas
}
