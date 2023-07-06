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
func (ss *SongService) AddSongInfo(song *model.SongInfo) (int64, int64, []model.SongInfo, error) {
	totals, currentPage, songlist, err := ss.songdao.AddSong(song)
	return totals, currentPage, songlist, err
}

// 获取十首歌曲
func (s *SongService) GetTenSongs() []model.SongInfo {
	songs := s.songdao.GetTenSongs()
	return songs
}

// 根据专辑id获取歌曲
func (ss *SongService) GetSongByAlbumid(album_id int) ([]model.SongInfo, error) {
	songlist, err := ss.songdao.GetSongByAlbumid(album_id)
	return songlist, err
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

	if l == 0 {
		return []model.SongInfo{}, err, nil, 0
	}

	if currentpage > pagenum {
		err1 := errors.New("当前要获取的歌曲分页过大！")
		return nil, err, err1, pagenum
	}

	var songpage []model.SongInfo

	if currentpage == 0 && pagenum == 1 {
		songpage = songs[0:l]
		return songpage, err, nil, pagenum
	}

	if currentpage == (pagenum-1) && remainder_flag {
		songpage = songs[(currentpage-1)*pagesize : l]
		return songpage, err, nil, pagenum
	}

	songpage = songs[currentpage*pagesize : (currentpage+1)*pagesize]
	return songpage, err, nil, pagenum
}

// 特定歌曲信息获取
func (ss *SongService) SongInfo(Name string) ([]model.SongInfo, error) {
	song, err := ss.songdao.GetSongbyName(Name)
	return song, err
}

// 特定页所有歌曲信息获取
func (ss *SongService) AllSongInfo(page int, pagesize int) ([]model.SongInfo, int64) {
	songlist, totalPage := ss.songdao.GetAllSongInfo(page, pagesize)
	return songlist, totalPage
}

// 歌曲信息修改
func (ss *SongService) ModifySongInfo(song *model.SongInfo) error {
	err := ss.songdao.ModifySong(song)
	return err
}

// 删除歌曲信息
func (ss *SongService) DeleteSongInfo(songID int) error {
	err := ss.songdao.DeleteSong(songID)
	return err
}

// 根据歌手id获取歌曲
func (as *SongService) GetSongByArtistid(artist_id int) ([]model.SongInfo, error) {
	songlist, err := as.songdao.GetSongByArtistid(artist_id)
	return songlist, err
}
