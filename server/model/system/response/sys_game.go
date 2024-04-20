package response

import "github.com/lindocedskes/model/system"

type GetGameListResponse struct {
	GameList []system.Question `json:"gameList"`
	//gameType := make(map[string]bool) //存储题目类型
	GameType map[string]bool                   `json:"gameType"`
	AllFiles []system.SysFileUploadAndDownload `json:"allFiles"`
}
