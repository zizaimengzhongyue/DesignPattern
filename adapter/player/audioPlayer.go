package player

import "fmt"

//AudioPlayer 音频播放器
type AudioPlayer interface {
	PlayMp3()
}

//Apple 苹果音频播放器
type Apple struct{}

//PlayMp3 播放 mp3
func (a Apple) PlayMp3() {
	fmt.Println("apple: 播放 mp3")
}
