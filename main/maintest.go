package main

import (
	"music-player/musicplayerserver/dao"
)

func main() {
	dao.Init()
	sd := dao.SongDao{}
	sd.CreateUserTable()
}
