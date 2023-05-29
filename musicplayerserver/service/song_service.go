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
func (s *SongService) GetSongLyric(name string) string {
	lyric := s.songdao.GetSongLr(name)
	if lyric != "" {
		fmt.Println("成功获取歌曲的歌词！")
	} else {
		fmt.Println("获取歌曲歌词失败！")
	}
	return lyric

}

// 获取歌曲详情
func (s *SongService) GetSongDetail(id string) (string, string, error) {
	songname, singer, err := s.songdao.GetSongInfo(id)
	return songname, singer, err
}
