package dao

import (
	"errors"
	"music-player/musicplayerserver/model"
)

type ArtistDao struct {
}

// 获取十个歌手信息
func (a *ArtistDao) GetTenArtist() []model.ArtistInfo {
	var artist []model.ArtistInfo
	var artists []model.ArtistInfo
	DB.Find(&artist)
	artists = artist[:10]
	return artists
}

// 根据歌手id获取歌手信息
func (a *ArtistDao) GetInfoById(id int) (model.ArtistInfo, error) {
	var artist model.ArtistInfo
	result := DB.First(&artist, id)
	return artist, result.Error
}

// 获取歌手描述
func (a *ArtistDao) GetProfile(id int) (string, error) {
	var artist model.ArtistInfo
	result := DB.First(&artist, id)
	return artist.Profile, result.Error
}

// 根据关键词获取歌手
func (a *ArtistDao) GetArtistByKeyWord(keyword string) ([]model.ArtistInfo, error) {
	var artist []model.ArtistInfo
	keyword = "%" + keyword + "%"
	result := DB.Where("name  LIKE ?", keyword).Find(&artist)
	return artist, result.Error
}

// 根据条件筛选歌手
func (a *ArtistDao) GetArtistByCondition(firstletter string, gender int, location int) ([]model.ArtistInfo, error) {
	var artist []model.ArtistInfo

	if firstletter == "0" && gender == 0 && location == 0 {
		result := DB.Find(&artist)
		return artist, result.Error
	} else if firstletter == "0" && gender == 0 {
		result := DB.Where("location = ?", location).Find(&artist)
		return artist, result.Error
	} else if firstletter == "0" && location == 0 {
		result := DB.Where("gender = ?", gender).Find(&artist)
		return artist, result.Error
	} else if gender == 0 && location == 0 {
		result := DB.Where("first_letter = ?", firstletter).Find(&artist)
		return artist, result.Error
	} else if firstletter == "0" {
		result := DB.Where("gender = ? AND location = ?", gender, location).Find(&artist)
		return artist, result.Error
	} else if gender == 0 {
		result := DB.Where("first_letter = ? AND location = ?", firstletter, location).Find(&artist)
		return artist, result.Error
	} else if location == 0 {
		result := DB.Where("first_letter = ? AND gender = ?", firstletter, gender).Find(&artist)
		return artist, result.Error
	}

	result := DB.Where("first_letter = ? AND gender = ? AND location = ?", firstletter, gender, location).Find(&artist)
	return artist, result.Error
}

// 返回特定页歌手信息
func (*ArtistDao) GetAllArtistInfo(page int, pagesize int) ([]model.ArtistInfo, int64) {
	var artistlist []model.ArtistInfo
	var totalrecord int64
	offset := (page - 1) * pagesize
	DB.Offset(offset).Limit(pagesize).Find(&artistlist).Offset(-1).Limit(-1).Count(&totalrecord)
	return artistlist, totalrecord
}

// 根据名字获取歌手信息
func (*ArtistDao) GetArtistbyName(name string) ([]model.ArtistInfo, error) {
	song := []model.ArtistInfo{}
	err := DB.Where(&song, "name=?", name).Find(&song).Error
	if err != nil {
		err = errors.New("查询歌曲信息出错！")
	}
	return song, err
}

// 添加歌手
func (*ArtistDao) AddArtist(artist *model.ArtistInfo) (int64, int64, []model.ArtistInfo, error) {
	var artistlist []model.ArtistInfo
	var totalrecord int64
	var offset int64
	var err error
	var currentPage int64
	DB.Create(artist)
	DB.Table("artist").Count(&totalrecord)
	if totalrecord%10 == 0 {
		offset = totalrecord - 10
		currentPage = totalrecord / 10
	} else {
		offset = totalrecord - (totalrecord % 10)
		currentPage = totalrecord/10 + 1
	}
	DB.Offset(int(offset)).Limit(10).Find(&artistlist)
	return totalrecord, currentPage, artistlist, err
}

// 删除歌手信息
func (*ArtistDao) DeleteArtist(artistID int) error {
	err := DB.Delete(&model.ArtistInfo{}, artistID).Error
	return err
}

// 修改歌手信息
func (*ArtistDao) ModifyArtist(artist *model.ArtistInfo) error {
	result := DB.Save(artist)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败！")
	}
	return err
}

// 根据专辑id获取特定歌手
func (s *ArtistDao) GetArtistByAlbumid(album_id int) ([]model.ArtistInfo, error) {

	var artist []model.ArtistInfo
	// 先在 album 数据库中根据 album_id 找到对应的 artist_id
	var artist_id int
	err := DB.Table("album").Select("artist_id").Where("album_id = ?", album_id).Row().Scan(&artist_id)
	if err != nil {
		return nil, err
	}

	// 在 artist 数据库中根据 artist_id 找到对应的艺术家信息
	result := DB.Table("artist").Where("artist_id = ?", artist_id).Find(&artist)
	return artist, result.Error
}
