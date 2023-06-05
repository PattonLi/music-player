package model

type SongInfo struct {
	Song_ID      int `gorm:"primarykey"`
	Album_ID     int
	Artist_ID    int
	Name         string
	Album        string
	Artist       string
	Pop          int
	Publish_time string
	Type         string
	Lyric_url    string
	Pic_url      string
	Duration     int
	Mark         int
	Url          string
	Lyric        string
}

// 获取表名，gorm创建表时会自己获取这个表名
func (s *SongInfo) TableName() string {
	return "song"
}
