package request

// 共用的请求结构体

// 存储主键id
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}
