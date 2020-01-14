//Description
//    没有想到面试场景下的例子
//    对一个播放器，开始设计的时候只能播放 mp3，在后续希望能够增加对 mp4 的支持，这时候已经存在 mp3 播放器和 mp4 播放器，解决方案是通过一个 adapter 将 mp3 播放器适配成 mp4 播放器
//    适配器这个地方用起来存在风险，因为在 advancedPlayer 看起来它是在播放一个 mp3，实际上是在播放 mp4
//    这里特别引用一点：适配器不是在详细设计时添加的，而是解决正在服役的项目的问题。
package main

import "./player"

func main() {
	advancedPlayer := player.AdvanceedPlayer{}

	mp3Player := player.Apple{}
	advancedPlayer.Play(mp3Player)

	mp4Player := player.Sony{}
	adapter := player.MediaPlayerAdapter{Player: mp4Player}
	advancedPlayer.Play(adapter)
}
