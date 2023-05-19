package main

import (
	musicplayerserver/dao"
)

func main() {
	dao.Init()
	sd := dao.SongDao{}
	sd.CreateUserTable()
}
