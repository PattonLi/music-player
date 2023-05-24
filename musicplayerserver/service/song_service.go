package service

import (
	"music-player/musicplayerserver/dao"
)

type SongService struct {
	songdao *dao.SongDao
}

// 构造函数
func NewSongService() *SongService {
	sgd := dao.NewSongDao()
	sg := SongService{
		songdao: sgd,
	}
	return &sg
}

// 获取音乐url
func (ss *SongService) GetSongURL(name string) string {
	song := ss.songdao.GetSongInfoByName(name)
	url := song.URL
	return url
}
