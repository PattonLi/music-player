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

func NewSongDao() *Songdao {
	return &Songdao{}
}
