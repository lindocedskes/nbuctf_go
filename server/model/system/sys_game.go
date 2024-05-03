package system

import "github.com/lindocedskes/model"

// 题目 model
type Question struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	IfHidden        bool              `json:"ifHidden" gorm:"default:true;comment:'是否隐藏'"`
	QueName         string            `json:"queName" gorm:"comment:'题目名称'"`               // 题目名称
	QueMark         int               `json:"queMark" gorm:"comment:'题目分数'"`               // 题目分数
	QueFlag         string            `json:"queFlag" gorm:"comment:'flag'"`               // flag
	QueDescribe     string            `json:"queDescribe" gorm:"type:text;comment:'题目描述'"` // 题目描述
	QueType         string            `json:"queType" gorm:"comment:'题目类型'"`               // 题目类型
	IfSolved        bool              `json:"ifSolved" gorm:"default:false;comment:'是否解决'"`
	QueSolvers      int               `json:"queSolvers" gorm:"default:0;comment:'解题者'"` // 已解题人数
	ImageUrl        string            `json:"imageUrl" gorm:"comment:'docker镜像地址'"`      // docker镜像地址
	InnerPort       string            `json:"innerPort" gorm:"comment:'docker内部访问端口'"`   // docker内部访问端口
	// 题目附件 多对多关系
	Files []SysFileUploadAndDownload `json:"files" gorm:"many2many:question_files;"`
}

func (Question) TableName() string {
	return "question"
}

type RightGameRecord struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	UserId          uint              `json:"userId" gorm:"comment:'用户ID'"`
	QuestionId      uint              `json:"questionId" gorm:"comment:'题目ID'"`
	SysUser         SysUser           `gorm:"foreignKey:UserId"`     // 声明UserId为外键
	Question        Question          `gorm:"foreignKey:QuestionId"` // 声明QuestionId为外键
}

type WrongGameRecord struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	SubmitFlag      string            `json:"submitFlag" gorm:"comment:'提交的flag'"`
	UserId          uint              `json:"userId" gorm:"comment:'用户ID'"`
	QuestionId      uint              `json:"questionId" gorm:"comment:'题目ID'"`
}

// 对应many2many:question_files
//type QuestionFile struct {
//	QuestionId uint `json:"questionId" gorm:"comment:'题目ID'"`
//	FileId     uint `json:"sys_file_upload_and_download_id" gorm:"comment:'文件ID'"`
//}

type UserScore struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	//主键
	UserId       uint   `json:"userId" gorm:"comment:'用户ID' primary_key"`
	UserName     string `json:"userName" gorm:"comment:'用户名'"`
	Score        int    `json:"score" gorm:"comment:'分数' default:0"`
	ScoreCrypto  int    `json:"scoreCrypto" gorm:"comment:'密码学分数' default:0"`
	ScoreWeb     int    `json:"scoreWeb" gorm:"comment:'web分数' default:0"`
	ScorePwn     int    `json:"scorePwn" gorm:"comment:'pwn分数' default:0"`
	ScoreMisc    int    `json:"scoreMisc" gorm:"comment:'misc分数' default:0"`
	ScoreReverse int    `json:"scoreReverse" gorm:"comment:'reverse分数' default:0"`
}
