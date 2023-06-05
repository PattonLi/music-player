package model

type AlbumInfo struct {
	AlbumID      int `gorm:"primarykey"`
	Size         int
	Name         string
	Artist       string
	Artist_ID    int
	Profile      string
	Type         string
	Publish_time string
	Pic_url      string
}

func (A *AlbumInfo) TableName() string {
	return "album"
}
