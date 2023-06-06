package dao

import (
	"errors"
	"math"
	"music-player/musicplayerserver/model"
)

type AlbumDao struct {
}

// 从数据库中获取十张专辑
func (a *AlbumDao) GetTenAlbums() []model.AlbumInfo {
	var album []model.AlbumInfo
	var albums []model.AlbumInfo
	DB.Find(&album)
	albums = album[:10]
	return albums
}

// 根据专辑id获取专辑
func (a *AlbumDao) GetAlbumById(id int) (model.AlbumInfo, error) {
	var album model.AlbumInfo
	result := DB.First(&album, id)

	return album, result.Error
}

// 获取特定专辑
func (*AlbumDao) GetAlbumbyName(name string) ([]model.AlbumInfo, error) {
	album := []model.AlbumInfo{}
	err := DB.Where(&album, "name=?", name).Find(&album).Error
	if err != nil {
		err = errors.New("查询专辑信息出错！")
	}
	return album, err
}

// 返回特定页所有专辑信息
func (*AlbumDao) GetAllAlbumInfo(page int, pagesize int) ([]model.AlbumInfo, int64) {
	var albumlist []model.AlbumInfo
	var totalrecord int64
	offset := (page - 1) * pagesize
	DB.Offset(offset).Limit(pagesize).Find(&albumlist).Offset(-1).Limit(-1).Count(&totalrecord)
	totalPage := int64(math.Ceil(float64(totalrecord) / float64(pagesize)))
	return albumlist, totalPage
// 根据歌手id获取歌手所有专辑
func (a *AlbumDao) GetAlbumByArtistid(artist_id int) ([]model.AlbumInfo, error) {
	var album []model.AlbumInfo
	result := DB.Where("artist_id = ?", artist_id).Find(&album)
	return album, result.Error
}
