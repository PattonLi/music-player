package model

type Playlist_Songs struct {
	Id          int
	Song_id     int
	Playlist_id int
}

func (p *Playlist_Songs) TableName() string {
	return "playlist_songs"
}
