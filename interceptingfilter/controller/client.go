package controller

//Client 客户端
type Client struct {
	FilterManager
}

//SendRequest 发送请求
//Params
//    request string 请求
func (c *Client) SendRequest(request string) {
	c.FilterManager.Execute(request)
}
