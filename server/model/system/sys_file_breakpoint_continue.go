package system

import "github.com/lindocedskes/model"

// file struct, 代表了一个完整的文件信息，以及关联的切片信息
type SysFile struct {
	model.BaseModel
	FileName     string
	FileMd5      string
	FilePath     string
	SysFileChunk []SysFileChunk //关联的文件切片
	ChunkTotal   int            //总切片数
	IsFinish     bool           //是否已经完成传输
}

// 代表了文件的一个切片，每个切片都会被单独存储在服务器上的某个位置
type SysFileChunk struct {
	model.BaseModel
	SysFileID       uint   //关联的文件ID（SysFileID）
	FileChunkNumber int    //切片的编号（FileChunkNumber
	FileChunkPath   string //切片的存储路径
}
