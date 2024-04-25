package response

import (
	"github.com/lindocedskes/model/system"
	"time"
)

type GetGameListResponse struct {
	GameList []system.Question `json:"gameList"`
	//gameType := make(map[string]bool) //存储题目类型
	GameType map[string]bool `json:"gameType"`
	//AllFiles []system.SysFileUploadAndDownload `json:"allFiles"`
}

// 提交时间 - 得分 折线图表
// 1. 根据RightGameRecord表 按用户id查询SysUser表得到用户名，按题目id查询到题目名称，按提交时间查询到提交时间
// 2. 根据Question表 按题目id查询到题目名称，按创建时间，题目分数
type ResSubmitScoreChart struct {
	UserId     uint      `json:"userId" gorm:"comment:'用户ID'"`        // 用户ID
	UserName   string    `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	QuestionId uint      `json:"questionId" gorm:"comment:'题目ID'"`    // 题目ID
	QueName    string    `json:"queName" gorm:"comment:'题目名称'"`       // 题目名称
	QueMark    int       `json:"queMark" gorm:"comment:'题目分数'"`       // 题目分数
	SubmitTime time.Time `json:"submitTime" gorm:"comment:'提交时间'"`    // 提交时间，升序
}
