package service

import (
	"errors"
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type ArtistService struct {
	artistdao *dao.ArtistDao
}

func NewArtistService() *ArtistService {
	a := dao.ArtistDao{}
	return &ArtistService{
		artistdao: &a,
	}
}

// 获取十个歌手信息
func (a *ArtistService) GetTenAritists() []model.ArtistInfo {
	artists := a.artistdao.GetTenArtist()
	return artists
}

// 根据歌手id获取歌手信息
func (a *ArtistService) GetArtistDetail(id int) (model.ArtistInfo, error) {
	artist, err := a.artistdao.GetInfoById(id)
	return artist, err
}

// 获得歌手描述
func (a *ArtistService) GetArtistDescribe(id int) (string, error) {
	profile, err := a.artistdao.GetProfile(id)
	return profile, err
}

// 根据关键词获取歌手
func (a *ArtistService) GetArtistByKeyWord(pagesize int, currentpage int, keyword string) ([]model.ArtistInfo, error, error, int) {
	artists, err := a.artistdao.GetArtistByKeyWord(keyword)
	l := len(artists)
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
		err1 := errors.New("当前要获取的歌手分页过大！")
		return nil, err, err1, pagenum
	}

	var artistpage []model.ArtistInfo

	if currentpage == pagenum && remainder_flag {
		artistpage = artists[(currentpage-1)*pagesize : l]
		return artistpage, err, nil, pagenum
	}

	artistpage = artists[(currentpage-1)*pagesize : currentpage*pagesize]
	return artistpage, err, nil, pagenum
}

// 获取对应筛选条件的歌手
func (a *ArtistService) GetAtrtistByContition(pagesize int, currentpage int, firstletter string, gender int, location int) ([]model.ArtistInfo, error, error, int) {
	artists, err := a.artistdao.GetArtistByCondition(firstletter, gender, location)
	l := len(artists)
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

	if l == 0 {
		return []model.ArtistInfo{}, err, nil, 0
	}

	if currentpage > pagenum {
		err1 := errors.New("当前要获取的歌手分页过大！")
		return []model.ArtistInfo{}, err, err1, pagenum
	}

	var artistpage []model.ArtistInfo

	if currentpage == 0 && pagenum == 1 {
		artistpage = artists[0:l]
		return artistpage, err, nil, pagenum
	}

	if currentpage == (pagenum-1) && remainder_flag {
		artistpage = artists[(currentpage-1)*pagesize : l]
		return artistpage, err, nil, pagenum
	}

	artistpage = artists[currentpage*pagesize : (currentpage+1)*pagesize]
	return artistpage, err, nil, pagenum

}

// 特定页歌手信息获取
func (as *ArtistService) AllArtistInfo(page int, pagesize int) ([]model.ArtistInfo, int64) {
	artistlist, totalrecord := as.artistdao.GetAllArtistInfo(page, pagesize)
	return artistlist, totalrecord
}

// 特定歌手信息获取
func (as *ArtistService) ArtistInfo(Name string) ([]model.ArtistInfo, error) {
	artist, err := as.artistdao.GetArtistbyName(Name)
	return artist, err
}

// 添加歌手信息
func (as *ArtistService) AddArtistInfo(artist *model.ArtistInfo) (int64, int64, []model.ArtistInfo, error) {
	totals, currentPage, artistlist, err := as.artistdao.AddArtist(artist)
	return totals, currentPage, artistlist, err
}

// 删除歌手信息
func (as *ArtistService) DeleteArtistInfo(artistID int) error {
	err := as.artistdao.DeleteArtist(artistID)
	return err
}

// 歌手信息修改
func (as *ArtistService) ModifyArtistInfo(artist *model.ArtistInfo) error {
	err := as.artistdao.ModifyArtist(artist)
	return err
}

// 根据专辑id获取歌手
func (as *ArtistService) GetArtistByAlbumid(album_id int) ([]model.ArtistInfo, error) {
	artistlist, err := as.artistdao.GetArtistByAlbumid(album_id)
	return artistlist, err
}
