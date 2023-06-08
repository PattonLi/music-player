package model

type LikesInfo struct {
	Like_id     int
	User_id     int
	Song_id     int
	Album_id    int
	Artist_id   int
	Playlist_id int
	Type        int
}

func (L *LikesInfo) TableName() string {
	return "likes"
}
