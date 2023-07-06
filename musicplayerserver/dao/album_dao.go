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
}

// 根据歌手id获取歌手所有专辑
func (a *AlbumDao) GetAlbumByArtistid(artist_id int) ([]model.AlbumInfo, error) {
	var album []model.AlbumInfo
	result := DB.Where("artist_id = ?", artist_id).Find(&album)
	return album, result.Error
}

// 根据关键词获取专辑
func (a *AlbumDao) GetAlbumByKeyWord(keyword string) ([]model.AlbumInfo, error) {
	var album []model.AlbumInfo
	keyword = "%" + keyword + "%"
	result := DB.Where("name  LIKE ?", keyword).Find(&album)
	return album, result.Error
}

// 添加专辑
func (*AlbumDao) AddAlbum(album *model.AlbumInfo) (int64, int64, []model.AlbumInfo, error) {
	var albumlist []model.AlbumInfo
	var totalrecord int64
	var offset int64
	var err error
	var currentPage int64
	DB.Create(album)
	DB.Table("album").Count(&totalrecord)
	if totalrecord%10 == 0 {
		offset = totalrecord - 10
		currentPage = totalrecord / 10
	} else {
		offset = totalrecord - (totalrecord % 10)
		currentPage = totalrecord/10 + 1
	}
	DB.Offset(int(offset)).Limit(10).Find(&albumlist)
	return totalrecord, currentPage, albumlist, err
}

// 修改专辑信息
func (*AlbumDao) ModifyAlbum(album *model.AlbumInfo) error {
	result := DB.Save(album)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败！")
	}
	return err
}

// 删除专辑信息
func (*AlbumDao) DeleteAlbum(albumID int) error {
	err := DB.Delete(&model.AlbumInfo{}, albumID).Error
	return err
}
