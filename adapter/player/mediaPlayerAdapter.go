package player

//MediaPlayerAdapter 视频播放适配器
type MediaPlayerAdapter struct {
	Player MediaPlayer
}

//PlayMp3 适配器将 mp4 播放适配为 mp3 播放
func (mpa MediaPlayerAdapter) PlayMp3() {
	mpa.Player.PlayMp4()
}
