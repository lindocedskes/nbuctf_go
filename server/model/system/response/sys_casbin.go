package response

import "github.com/lindocedskes/model/system/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"` //存储多个Casbin策略信息。
}
