package model

type Comments struct {
	Nickname    string `json:"nickname"`
	Picurl      string `json:"picUrl"`
	Comment     string `json:"comment"`
	Commentid   int    `json:"commentId"`
	Commenttime string `json:"commentTime"`
	Like        int    `json:"like"`
}
