package utils

// 规则集合
var (
	LoginVerify       = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AuthorityIdVerify = Rules{"AuthorityId": {NotEmpty()}} // cabin角色id权限校验
)
