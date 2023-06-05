package model

type ArtistInfo struct {
	ArtistID   int    `json:"artistId" gorm:"primarykey"`
	Name       string `json:"artist"`
	Profile    string `json:"profile"`
	Location   string `json:"location"`
	Song_size  int    `json:"songSize"`
	Album_size int    `json:"albumSize"`
	Pic_url    string `json:"picUrl"`
}

func (a *ArtistInfo) TableName() string {
	return "artist"
}
