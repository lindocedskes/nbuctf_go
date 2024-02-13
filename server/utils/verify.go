package utils

// 规则集合
var (
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}} // cabin角色id权限校验
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	IdVerify               = Rules{"ID": []string{NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
)
