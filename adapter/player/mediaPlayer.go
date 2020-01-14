package player

import "fmt"

//MediaPlayer 视频播放器
type MediaPlayer interface {
	PlayMp4()
}

//Sony sony 播放器
type Sony struct{}

//PlayMp4 播放 mp4
func (s Sony) PlayMp4() {
	fmt.Println("sony: 播放 mp4")
}
