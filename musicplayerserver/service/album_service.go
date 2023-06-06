package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type AlbumService struct {
	albumdao *dao.AlbumDao
}

func NewAlbumService() *AlbumService {
	a := dao.AlbumDao{}
	al := AlbumService{
		albumdao: &a,
	}
	return &al
}

// 获取十张专辑
func (al *AlbumService) GetTenAlbums() []model.AlbumInfo {
	albums := al.albumdao.GetTenAlbums()
	return albums
}

// 根据专辑id获取专辑
func (al *AlbumService) GetAlbumById(id int) (model.AlbumInfo, error) {
	album, err := al.albumdao.GetAlbumById(id)
	return album, err
}

// 特定专辑信息获取
func (as *AlbumService) AlbumInfo(Name string) ([]model.AlbumInfo, error) {
	album, err := as.albumdao.GetAlbumbyName(Name)
	return album, err
}

// 特定页所有专辑信息获取
func (als *AlbumService) AllAlbumInfo(page int, pagesize int) ([]model.AlbumInfo, int64) {
	albumlist, totalPage := als.albumdao.GetAllAlbumInfo(page, pagesize)
	return albumlist, totalPage
}
