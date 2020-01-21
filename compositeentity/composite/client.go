package composite

import "fmt"

//Client 客户端
type Client struct {
	Composite
}

//SetData 写入数据
//Params
//    data1 数据1
//    data2 数据2
func (c *Client) SetData(data1 string, data2 string) {
	c.Composite.SetData(data1, data2)
}

//GetData 读取数据
func (c *Client) GetData() {
	datas := c.Composite.GetData()
	fmt.Println(datas)
}
