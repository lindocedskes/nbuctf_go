package response

import "github.com/lindocedskes/model/system"

type FileResponse struct {
	File system.SysFile `json:"file"`
}

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}
