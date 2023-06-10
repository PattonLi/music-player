package model

type MvInfo struct {
	Movie_id int    `json:"movieId"`
	Url      string `json:"url"`
	Pic_url  string `json:"picUrl"`
	Duration int    `json:"duration"`
	Movie    string `json:"movie"`
	Artist   string `json:"artist"`
}

func (m *MvInfo) TableName() string {
	return "mv"
}
