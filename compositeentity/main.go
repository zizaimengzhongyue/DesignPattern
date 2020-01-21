//Description
//    看起来就是一层一层的打包
package main

import "./composite"

func main() {
	client := composite.Client{}
	client.SetData("hello", "world")
	client.GetData()
	client.SetData("first", "second")
	client.GetData()
}
