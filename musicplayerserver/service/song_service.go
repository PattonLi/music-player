package service

import (
	"fmt"
	"music-player/musicplayerserver/dao"
)

type SongService struct {
	songservice *dao.Songdao
}

// 构造函数
func (*SongService) NewSongService() *SongService {
	s := dao.Songdao{}
	sv := SongService{
		songservice: &s,
	}
	return &sv
}

// 获取歌曲信息
func (s *SongService) GetSong(id string) string {
	lyric := s.songservice.GetSongInfo(id)
	if lyric != "" {
		fmt.Println("成功获取歌曲的信息！")
	} else {
		fmt.Println("获取歌曲信息失败！")
	}
	return lyric
}
