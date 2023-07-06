package model

type LogInfo struct {
	Log_id   int	`json:"log_id"`
	User_id  int	`json:"user_id"`
	Song_id  int	`json:"song_id"`
	Songname string	`json:"songname"`
	Username string	`json:"username"`
	Time     string	`json:"time"`
	Type     int	`json:"type"`
}

func (l *LogInfo) TableName() string {
	return "log"
}
