package service

import (
	"errors"
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

// 获取十首歌曲
func (s *SongService) GetTenSongs() []model.SongInfo {
	songs := s.songdao.GetTenSongs()
	return songs
}

// 获取专辑的所有歌曲
func (s *SongService) GetAlbumSongs(albumid int) ([]model.SongInfo, error) {
	songs, err := s.songdao.GetSongsInAlbum(albumid)
	return songs, err
}

// 根据时间或者热度对歌手歌曲进行排序,获取歌曲分页
func (s *SongService) GetSongsPage(id int, order string, currentpage int, pagesize int) ([]model.SongInfo, error, error, int) {
	songs, err := s.songdao.SortSongsByOrder(id, order)
	l := len(songs)
	n := l / pagesize
	remainder := l % pagesize
	var pagenum int
	var remainder_flag bool
	if remainder == 0 {
		pagenum = n
		remainder_flag = false
	} else {
		pagenum = n + 1
		remainder_flag = true
	}

	if currentpage > pagenum {
		err1 := errors.New("当前要获取的歌曲分页过大！")
		return nil, err, err1, pagenum
	}

	var songpage []model.SongInfo

	if currentpage == pagenum && remainder_flag {
		songpage = songs[(currentpage-1)*pagesize : l]
		return songpage, err, nil, pagenum
	}

	songpage = songs[(currentpage-1)*pagesize : currentpage*pagesize]
	return songpage, err, nil, pagenum
}

// 根据songid获取歌曲
func (s *SongService) GetSongDetail(id int) (model.SongInfo, error) {
	song, err := s.songdao.GetSongDetail(id)
	return song, err
}

// 根据关键词获取歌曲
func (s *SongService) GetSongByKeyWord(pagesize int, currentpage int, keyword string) ([]model.SongInfo, error, error, int) {
	songs, err := s.songdao.GetSongByKeyWord(keyword)
	l := len(songs)
	n := l / pagesize
	remainder := l % pagesize
	var pagenum int
	var remainder_flag bool
	if remainder == 0 {
		pagenum = n
		remainder_flag = false
	} else {
		pagenum = n + 1
		remainder_flag = true
	}

	if currentpage > pagenum {
		err1 := errors.New("当前要获取的歌曲分页过大！")
		return nil, err, err1, pagenum
	}

	var songpage []model.SongInfo

	if currentpage == pagenum && remainder_flag {
		songpage = songs[(currentpage-1)*pagesize : l]
		return songpage, err, nil, pagenum
	}

	songpage = songs[(currentpage-1)*pagesize : (currentpage)*pagesize]
	return songpage, err, nil, pagenum
}
