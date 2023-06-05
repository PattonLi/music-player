package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
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
func (s *SongService) GetSongLyric(id string) (string, error) {
	lyric, err := s.songdao.GetSongLr(id)
	return lyric, err

}

// 获取歌曲详情
/*func (s *SongService) GetSongDetail(id string) (string, string, error) {
	songname, singer, err := s.songdao.GetSongInfo(id)
	return songname, singer, err
}*/

// 添加歌曲
func (s *SongService) AddSong(song model.SongInfo) bool {
	result := s.songdao.AddSong(&song)
	return result
}
