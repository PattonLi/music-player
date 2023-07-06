package dao

import (
	"errors"
	"music-player/musicplayerserver/model"

	"gorm.io/gorm"
)

type AlbumDao struct {
}

// 从数据库中获取十张专辑
func (a *AlbumDao) GetTenAlbums() []model.AlbumInfo {
	var album []model.AlbumInfo
	var albums []model.AlbumInfo
	DB.Find(&album)
	albums = album[:30]
	return albums
}

// 根据专辑id获取专辑
func (a *AlbumDao) GetAlbumById(id int) (model.AlbumInfo, error) {
	var album model.AlbumInfo
	result := DB.First(&album, id)

	return album, result.Error
}

// 获取特定专辑
func (dao *AlbumDao) GetAlbumbyName(name string) ([]model.AlbumInfo, error) {
	album := []model.AlbumInfo{}
	err := DB.Where("name LIKE ?", "%"+name+"%").Find(&album).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查找不到专辑信息！")
	} else {
		err = nil
	}
	return album, err
}

// 返回特定页所有专辑信息
func (*AlbumDao) GetAllAlbumInfo(page int, pagesize int) ([]model.AlbumInfo, int64) {
	var albumlist []model.AlbumInfo
	var totalrecord int64
	offset := (page - 1) * pagesize
	DB.Offset(offset).Limit(pagesize).Find(&albumlist).Offset(-1).Limit(-1).Count(&totalrecord)
	return albumlist, totalrecord
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
	// var err error = nil
	if result.Error != nil {
		return errors.New("修改失败！")
	}
	return nil
}

// 删除专辑信息
func (*AlbumDao) DeleteAlbum(albumID int) error {
	err := DB.Delete(&model.AlbumInfo{}, albumID).Error
	return err
}
