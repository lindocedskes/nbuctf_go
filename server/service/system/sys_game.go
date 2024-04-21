package system

import (
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GameService struct{}

// 获取所有未隐藏题目列表
func (gameService *GameService) GetGameList() ([]system.Question, error) {
	db := global.NBUCTF_DB.Model(&system.Question{})
	var gameList []system.Question
	err := db.Preload("Files").Find(&gameList, "if_hidden = ?", false).Error
	return gameList, err
}

// 判断当前用户是否已解决该题
func (gameService *GameService) IfSolved(queID uint, userID uint) bool {
	db := global.NBUCTF_DB.Model(&system.RightGameRecord{})
	var gameRecord system.RightGameRecord
	err := db.Where("question_id = ? AND user_id = ?", queID, userID).First(&gameRecord).Error
	if err != nil {
		return false
	}
	return true
}

// 获取所有文件
func (gameService *GameService) GetAllFiles() ([]system.SysFileUploadAndDownload, error) {
	db := global.NBUCTF_DB.Model(&system.SysFileUploadAndDownload{})
	var allFiles []system.SysFileUploadAndDownload
	err := db.Find(&allFiles).Error
	return allFiles, err
}

// 提交flag，记录，正确则加分
func (gameService *GameService) JudgeFlag(submitFlag string, queID uint, userID uint) (bool, error) {
	var question system.Question
	global.NBUCTF_DB.Model(&system.Question{}).Where("id = ?", queID).First(&question) //查询题目
	if question.QueFlag == submitFlag {

		var rightRecord system.RightGameRecord
		rightRecord.QuestionId = queID
		rightRecord.UserId = userID
		err := global.NBUCTF_DB.Create(&rightRecord).Error //创建正确题目记录
		if err != nil {
			return false, err
		}
		//用户加分
		var userScore system.UserScore
		var userName string
		_ = global.NBUCTF_DB.Model(&system.SysUser{}).Select("username").Where("id = ?", userID).First(&userName)
		if err := global.NBUCTF_DB.Model(&system.UserScore{}).Where("user_id = ?", userID).First(&userScore).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 用户记录不存在，创建新记录并设置初始分数为0
				userScore = system.UserScore{
					UserId:   userID,
					UserName: userName,
					Score:    0,
				}
				if err := global.NBUCTF_DB.Create(&userScore).Error; err != nil {
					return false, err
				}
			} else {
				return false, err
			}
		}
		userScore.Score += question.QueMark
		err = global.NBUCTF_DB.Model(&system.UserScore{}).Where("user_id = ?", userID).Updates(system.UserScore{Score: userScore.Score}).Error //更新用户分数，按UserID
		if err != nil {
			global.GVA_LOG.Error("分数更新失败!", zap.Error(err))
			return false, err
		}
		return true, nil

	} else {
		var wrongRecord system.WrongGameRecord
		wrongRecord.QuestionId = queID
		wrongRecord.UserId = userID
		wrongRecord.SubmitFlag = submitFlag
		err := global.NBUCTF_DB.Create(&wrongRecord).Error //创建错误题目记录
		if err != nil {
			return false, err
		}
		return false, nil
	}
}

// 获取排行榜
func (gameService *GameService) GetRankList() ([]system.UserScore, error) {
	db := global.NBUCTF_DB.Model(&system.UserScore{})
	var rankList []system.UserScore
	err := db.Order("score desc, updated_at asc").Find(&rankList).Error //按分数降序排列，相同分数按更新时间升序
	return rankList, err
}

// 获取题目附件下载 by question_id
func (gameService *GameService) GetFileById(queID uint) ([]system.SysFileUploadAndDownload, error) {
	//	Files []SysFileUploadAndDownload `json:"files" gorm:"many2many:question_files;"`
	var question system.Question
	err := global.NBUCTF_DB.Preload("Files").First(&question, queID).Error
	//再根据file_id查找对应的文件
	files := question.Files
	return files, err
}

// 获取所有题目列表 admin
func (gameService *GameService) GetGameListByAdmin() ([]system.Question, error) {
	db := global.NBUCTF_DB.Model(&system.Question{})
	var gameList []system.Question
	//获取所有题目列表
	err := db.Preload("Files").Find(&gameList).Error
	return gameList, err
}

// 编辑题目，更新题目信息 by id
func (gameService *GameService) EditQuestion(question system.Question) error {
	err := global.NBUCTF_DB.Model(&system.Question{}).Where("id = ?", question.ID).Updates(&question).Error
	return err
}

// 创建题目
func (gameService *GameService) CreateQuestion(question system.Question) (system.Question, error) {
	err := global.NBUCTF_DB.Create(&question).Error
	return question, err
}

// 删除题目 by id，并删除与之关联的file
func (gameService *GameService) DeleteQuestion(queID uint) error {
	var question system.Question
	err := global.NBUCTF_DB.First(&question, queID).Error
	if err != nil {
		return err
	}

	// 删除与 Question 关联的记录
	global.NBUCTF_DB.Model(&question).Association("Files").Clear()

	// 删除 Question
	err = global.NBUCTF_DB.Delete(&question).Error
	if err != nil {
		return err
	}

	return nil
}

// 添加题目附件 gorm Association，Append
func (gameService *GameService) AddFile(queID uint, fileID uint) error {
	var question system.Question
	err := global.NBUCTF_DB.First(&question, queID).Error
	if err != nil {
		return err
	}

	var file system.SysFileUploadAndDownload
	err = global.NBUCTF_DB.First(&file, fileID).Error
	if err != nil {
		return err
	}

	// 添加文件到题目
	global.NBUCTF_DB.Model(&question).Association("Files").Append(&file)
	return nil
}
