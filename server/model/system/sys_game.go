package system

import "github.com/lindocedskes/model"

// 题目 model
type Question struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	IfHidden        bool              `json:"ifHidden" gorm:"default:false;comment:'是否隐藏'"`
	QueName         string            `json:"queName" gorm:"comment:'题目名称'"`     // 题目名称
	QueMark         int               `json:"queMark" gorm:"comment:'题目分数'"`     // 题目分数
	QueFlag         string            `json:"queFlag" gorm:"comment:'flag'"`     // flag
	QueDescribe     string            `json:"queDescribe" gorm:"comment:'题目描述'"` // 题目描述
	QueType         string            `json:"queType" gorm:"comment:'题目类型'"`     // 题目类型
	IfSolved        bool              `json:"ifSolved" gorm:"default:false;comment:'是否解决'"`
}

func (Question) TableName() string {
	return "question"
}

type RightGameRecord struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	UserId          uint              `json:"userId" gorm:"comment:'用户ID'"`
	QuestionId      uint              `json:"questionId" gorm:"comment:'题目ID'"`
}

type WrongGameRecord struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	SubmitFlag      string            `json:"submitFlag" gorm:"comment:'提交的flag'"`
	UserId          uint              `json:"userId" gorm:"comment:'用户ID'"`
	QuestionId      uint              `json:"questionId" gorm:"comment:'题目ID'"`
}

type QuestionFile struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	QuestionId      uint              `json:"questionId" gorm:"comment:'题目ID'"`
	FileId          uint              `json:"fileId" gorm:"comment:'文件ID'"`
}

type UserScore struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	//主键
	UserId   uint   `json:"userId" gorm:"comment:'用户ID' primary_key"`
	UserName string `json:"userName" gorm:"comment:'用户名'"`
	Score    int    `json:"score" gorm:"comment:'分数' default:0"`
}
