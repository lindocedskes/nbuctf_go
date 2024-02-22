package response

import "github.com/lindocedskes/model/system"

type SysFileResponse struct {
	File system.SysFileUploadAndDownload `json:"file"`
}
