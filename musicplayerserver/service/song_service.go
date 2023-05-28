package service

import (
	"fmt"
	"music-player/musicplayerserver/dao"
)

type SongService struct {
	songdao *dao.Songdao
}

// 构造函数
func (*SongService) NewSongService() *SongService {
	s := dao.Songdao{}
	sv := SongService{
		songdao: &s,
	}
	return &sv
}

// 获取歌曲歌词
func (s *SongService) GetSong(name string) string {
	lyric := s.songdao.GetSongInfo(name)
	if lyric != "" {
		fmt.Println("成功获取歌曲的歌词！")
	} else {
		fmt.Println("获取歌曲歌词失败！")
	}
	return lyric

}
