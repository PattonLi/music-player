package dao

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/model"
)

type Songdao struct {
}

// 查询歌曲歌词
func (s *Songdao) GetSongLr(name string) string {
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

// 获取歌曲详情
func (s *Songdao) GetSongInfo(id string) (string, string, error) {
	song := model.SongInfo{}
	err := DB.Find(&song, id).Error
	fmt.Print(id)
	if err != nil {
		err = errors.New("歌曲不存在！")
	}
	return song.Name, song.Singer, err
}

func NewSongDao() *Songdao {
	return &Songdao{}
}
