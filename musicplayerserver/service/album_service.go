package service

import (
	"errors"
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

// 分页获取歌手专辑
func (al *AlbumService) GetAlbumPage(artist_id int, currentpage int, pagesize int) ([]model.AlbumInfo, error, error, int) {
	albums, err := al.albumdao.GetAlbumByArtistid(artist_id)
	l := len(albums)
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
		err1 := errors.New("当前要获取的专辑分页过大！")
		return nil, err, err1, pagenum
	}

	var albumpage []model.AlbumInfo

	if currentpage == pagenum && remainder_flag {
		albumpage = albums[(currentpage-1)*pagesize : l]
		return albumpage, err, nil, pagenum
	}

	albumpage = albums[(currentpage-1)*pagesize : currentpage*pagesize]
	return albumpage, err, nil, pagenum
}
