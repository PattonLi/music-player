package dao

import (
	"errors"
	"math"
	"music-player/musicplayerserver/model"

	"gorm.io/gorm"
)

type Songdao struct {
}

// 查询歌曲歌词
func (s *Songdao) GetSongLr(id string) (string, error) {
	song := model.SongInfo{}
	//DB.Where("name = ?", name).Find(&song)
	err := DB.First(&song, id).Error
	return song.Lyric, err
}

// 添加歌曲
func (*Songdao) AddSong(song *model.SongInfo) (int64, int64, []model.SongInfo, error) {
	newsong := model.SongInfo{}
	var songlist []model.SongInfo
	var totalrecord int64
	var offset int64
	var err error
	var currentPage int64
	DB.Take(&newsong, "name = ?", song.Name)
	if newsong.Album_ID != 0 {
		err = errors.New("歌曲已添加！")
		return 0, 0, songlist, err
	}
	DB.Create(song)
	DB.Table("song").Count(&totalrecord)
	if totalrecord%10 == 0 {
		offset = totalrecord - 10
		currentPage = totalrecord / 10
	} else {
		offset = totalrecord - (totalrecord % 10)
		currentPage = totalrecord/10 + 1
	}
	DB.Offset(int(offset)).Limit(10).Find(&songlist)
	return totalrecord, currentPage, songlist, err
}

// 获取歌曲详情
/*func (s *Songdao) GetSongInfo(id string) (string, string, error) {
	song := model.SongInfo{}
	err := DB.First(&song, id).Error

}*/

// 从数据库中获取十首歌曲
func (s *Songdao) GetTenSongs() []model.SongInfo {
	var song []model.SongInfo
	var songs []model.SongInfo
	count := 0
	DB.Find(&song)

	for i := 0; i < len(song); i += 10 {
		songs = append(songs, song[i])
		count++

		if count == 30 {
			break
		}
	}
	return songs
}

// 根据专辑id获取歌曲
func (s *Songdao) GetSongByAlbumid(album_id int) ([]model.SongInfo, error) {
	var song []model.SongInfo
	result := DB.Where("album_id = ?", album_id).Find(&song)
	return song, result.Error
}

func NewSongDao() *Songdao {
	return &Songdao{}
}

// 根据热度或者时间对歌手歌曲进行排序
func (s *Songdao) SortSongsByOrder(id int, order string) ([]model.SongInfo, error) {
	var song []model.SongInfo

	if order == "hot" {
		order = "pop"
	}
	result := DB.Order(order).Where("artist_id = ?", id).Find(&song)
	return song, result.Error
}

// 根据歌曲id获得歌曲
func (s *Songdao) GetSongDetail(id int) (model.SongInfo, error) {
	var song model.SongInfo
	result := DB.First(&song, id)
	return song, result.Error
}

// 根据关键词获取歌曲
func (s *Songdao) GetSongByKeyWord(keyword string) ([]model.SongInfo, error) {
	var song []model.SongInfo
	keyword = "%" + keyword + "%"
	result := DB.Where("name  LIKE ?", keyword).Find(&song)
	return song, result.Error
}

// 根据名字获取歌曲信息
func (*Songdao) GetSongbyName(name string) ([]model.SongInfo, error) {
	song := []model.SongInfo{}
	err := DB.Where("name LIKE ?", "%"+name+"%").Find(&song).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || DB.RowsAffected == 0 {
		err = errors.New("查找不到歌曲信息！")
	}
	return song, err
}

// 返回特定页所有歌曲信息
func (*Songdao) GetAllSongInfo(page int, pagesize int) ([]model.SongInfo, int64) {
	var songlist []model.SongInfo
	var totalrecord int64
	offset := (page - 1) * pagesize
	DB.Offset(offset).Limit(pagesize).Find(&songlist).Offset(-1).Limit(-1).Count(&totalrecord)
	totalPage := int64(math.Ceil(float64(totalrecord) / float64(pagesize)))
	return songlist, totalPage
}

// 修改歌曲信息
func (*Songdao) ModifySong(song *model.SongInfo) error {
	result := DB.Save(song)
	var err error = nil
	if result.Error != nil {
		err = errors.New("修改失败！")
	}
	return err
}

// 删除歌曲信息
func (*Songdao) DeleteSong(songID int) error {
	err := DB.Delete(&model.SongInfo{}, songID).Error
	return err
}

// 根据歌手id获取歌曲
func (a *Songdao) GetSongByArtistid(artist_id int) ([]model.SongInfo, error) {
	var song []model.SongInfo
	result := DB.Where("artist_id = ?", artist_id).Find(&song)
	return song, result.Error
}
