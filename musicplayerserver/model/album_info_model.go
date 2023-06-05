package model

type AlbumInfo struct {
	AlbumID      int    `json:"albumId" gorm:"primarykey"`
	Size         int    `json:"size"`
	Name         string `json:"name"`
	Artist       string `json:"artist"`
	Artist_ID    int    `json:"artistId"`
	Profile      string `json:"profile"`
	Type         string `json:"type"`
	Publish_time string `json:"publishTime"`
	Pic_url      string `json:"picUrl"`
}

func (A *AlbumInfo) TableName() string {
	return "album"
}
