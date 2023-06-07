package dao

import (
	"music-player/musicplayerserver/model"
)

type Songdao struct {
}

// 查询歌曲歌词
func (s *Songdao) GetSongLr(id string) (string, error) {
	song := model.SongInfo{}
	//DB.Where("name = ?", name).Find(&song)
	err := DB.First(&song, id).Error
	return song.Lyric, err
}

// 添加歌曲
func (s *Songdao) AddSong(song *model.SongInfo) bool {
	result := DB.Create(song)
	if result.Error != nil {
		return false
	} else {
		return true
	}
}

// 获取歌曲详情
/*func (s *Songdao) GetSongInfo(id string) (string, string, error) {
	song := model.SongInfo{}
	err := DB.First(&song, id).Error

}*/

// 从数据库中获取十首歌曲
func (s *Songdao) GetTenSongs() []model.SongInfo {
	var song []model.SongInfo
	var songs []model.SongInfo
	DB.Find(&song)
	songs = song[:10]
	return songs
}

// 获取专辑中的所有歌曲
func (s *Songdao) GetSongsInAlbum(albumid int) ([]model.SongInfo, error) {
	var song []model.SongInfo
	result := DB.Where("album_id = ?", albumid).Find(&song)
	return song, result.Error
}

func NewSongDao() *Songdao {
	return &Songdao{}
}

// 根据热度或者时间对歌手歌曲进行排序
func (s *Songdao) SortSongsByOrder(id int, order string) ([]model.SongInfo, error) {
	var song []model.SongInfo
	result := DB.Order(order).Where("artist_id = ?", id).Find(&song)
	return song, result.Error
}

// 根据歌曲id获得歌曲
func (s *Songdao) GetSongDetail(id int) (model.SongInfo, error) {
	var song model.SongInfo
	result := DB.First(&song, id)
	return song, result.Error
}

// 根据关键词获取歌曲
func (s *Songdao) GetSongByKeyWord(keyword string) ([]model.SongInfo, error) {
	var song []model.SongInfo
	keyword = "%" + keyword + "%"
	result := DB.Where("name  LIKE ?", keyword).Find(&song)
	return song, result.Error
}
