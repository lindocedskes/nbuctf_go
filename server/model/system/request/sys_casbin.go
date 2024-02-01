package request

// 请求中的策略信息
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// 权限id，【路径，方法】 与数据表的策略进行匹配验证
type CasbinRequestReceive struct {
	AuthorityId uint         `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"` //存储多个Casbin策略信息。
}
