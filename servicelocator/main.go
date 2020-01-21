//Description
//    缓存，我以前在某个 PHP 项目中看到过类似的方式，为了提高文件加载速度，会在每次查找文件的时候把文件的路径缓存下来
package main

import "./controller"

func main() {
	locator := &controller.Locator{}
	con01 := locator.GetController("controller01")
	con01.Handle("hello")
	con02 := locator.GetController("controller02")
	con02.Handle("hello")
}
