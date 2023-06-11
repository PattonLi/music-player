package model

type LogInfo struct {
	Log_id   int
	User_id  int
	Song_id  int
	Songname string
	Username string
	Time     string
	Type     int
}

func (l *LogInfo) TableName() string {
	return "log"
}
