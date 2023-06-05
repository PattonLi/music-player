package model

type SongInfo struct {
	Song_ID      int    `json:"songId" gorm:"primarykey"`
	Album_ID     int    `json:"albumId"`
	Artist_ID    int    `json:"artistId"`
	Name         string `json:"name"`
	Album        string `json:"album"`
	Artist       string `json:"artist"`
	Pop          int    `json:"pop"`
	Publish_time string `json:"publishTime"`
	Type         string `json:"type"`
	Lyric_url    string `json:"lyricUrl"`
	Pic_url      string `json:"picUrl"`
	Duration     int    `json:"duration"`
	Mark         int    `json:"mark"`
	Url          string `json:"url"`
	Lyric        string `json:"lyric"`
}

// 获取表名，gorm创建表时会自己获取这个表名
func (s *SongInfo) TableName() string {
	return "song"
}
