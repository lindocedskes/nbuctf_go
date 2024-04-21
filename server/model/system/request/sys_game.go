package request

// 对应many2many:question_files
type QuestionFile struct {
	QuestionId uint `json:"questionId" gorm:"comment:'题目ID'"`
	FileId     uint `json:"fileId" gorm:"comment:'文件ID'"`
}
