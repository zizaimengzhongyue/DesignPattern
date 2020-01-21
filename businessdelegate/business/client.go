package business

//Client 客户端
type Client struct {
	DeleGate *DeleGate
}

//SetDeleGate 设置代表
//Params
//    d *DeleGate 代表
func (c *Client) SetDeleGate(d *DeleGate) {
	c.DeleGate = d
}

//DoTask 执行任务
func (c *Client) DoTask() {
	c.DeleGate.DoTask()
}
