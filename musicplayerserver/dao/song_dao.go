package dao

import (
	"fmt"
	"music-player/musicplayerserver/model"
)

type SongDao struct {
}

func NewSongDao() *SongDao {
	return &SongDao{}
}

// 创建表，在第一次运行创建就行了，后面不用管他
func (*SongDao) CreateUserTable() {
	err := DB.AutoMigrate(&model.Song{})
	if err != nil {
		fmt.Println("Create Song Table failed: ", err)
	} else {
		fmt.Println("Successfully create Song Table.")
	}
}

// 通过id查询音乐信息
func (*SongDao) GetSongInfoByID(id string) *model.Song {
	user := model.Song{}
	DB.First(&user, id)
	return &user
}

// 通过音乐名查询音乐信息
func (*SongDao) GetSongInfoByName(name string) *model.Song {
	song := model.Song{}
	DB.Where("name = ?", name).First(&song)
	return &song
}

// 添加音乐
func (*SongDao) AddSong(song *model.Song) bool {
	result := DB.Create(song)
	if result.Error != nil {
		panic("Insert song error")
	} else {
		fmt.Println("Successfully insert song.")
		return true
	}
}
