package model

type PlaylistInfo struct {
	Playlist_id int    `json:"playListId"`
	User_id     int    `json:"userId"`
	Mark        int    `json:"mark"`
	Size        int    `json:"size"`
	Create_time string `json:"createTime"`
	User        string `json:"user"`
	Playlist    string `json:"playList"`
	Profile     string `json:"profile"`
	Tag         string `json:"tag"`
	Pic_url     string `json:"picUrl"`
}

func (p *PlaylistInfo) TableName() string {
	return "playlist"
}
