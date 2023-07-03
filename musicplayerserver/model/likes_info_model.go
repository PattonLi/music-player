package model

type LikesInfo struct {
	Like_id     int
	User_id     int `json:"userId"`
	Song_id     int `json:"songId"`
	Album_id    int `json:"albumId"`
	Artist_id   int `json:"artistId"`
	Playlist_id int `json:"playlistId"`
	Type        int `json:"type"`
}

func (L *LikesInfo) TableName() string {
	return "likes"
}
