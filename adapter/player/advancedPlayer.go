package player

//AdvanceedPlayer 播放器
type AdvanceedPlayer struct{}

//Play 播放
func (ap AdvanceedPlayer) Play(a AudioPlayer) {
	a.PlayMp3()
}
