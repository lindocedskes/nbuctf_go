package request

// 对应many2many:question_files
type QuestionFile struct {
	QuestionId uint `json:"questionId" gorm:"comment:'题目ID'"`
	FileId     uint `json:"fileId" gorm:"comment:'文件ID'"`
}

// 排行榜
type RankListByType struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
	Type     string `json:"type" form:"type"`         // 题目类型
}
