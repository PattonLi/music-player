package model

type CommentInfo struct {
	Comment_id     int    `json:"commentId" gorm:"primarykey"`
	Comment        string `json:"comment"`
	Comment_time   string `json:"commentTime"`
	Like           int    `json:"like"`
	Ref_comment_id int    `json:"refCommentID"`
	Song_id        int    `json:"songID"`
	Type           int    `json:"type"`
	User_id        int    `json:"userID"`
}

func (c *CommentInfo) TableName() string {
	return "comment"
}
