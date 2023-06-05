package model

type ArtistInfo struct {
	ArtistID   int `gorm:"primarykey"`
	Name       string
	Profile    string
	Location   string
	Song_size  int
	Album_size int
	Pic_url    string
}

func (a *ArtistInfo) TableName() string {
	return "artist"
}
