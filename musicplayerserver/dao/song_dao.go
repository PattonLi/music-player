package dao

import (
	"music-player/musicplayerserver/model"
)

type Songdao struct {
}

// 查询歌曲歌词
func (s *Songdao) GetSongInfo(name string) string {
	song := model.SongInfo{}
	DB.Where("name = ?", name).Find(&song)
	return song.Lyric
}

// 添加歌曲
/*func (s *Songdao) AddSong(song *model.SongInfo) bool {
	result := DB.Create(song)
	if result.Error != nil {
		panic("Insert error")
	} else {
		fmt.Println("Successfully insert.")
		return true
	}
}*/

func NewSongDao() *Songdao {
	return &Songdao{}
}
