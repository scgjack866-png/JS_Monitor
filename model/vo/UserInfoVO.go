package vo

type UserInfo struct {
	// 用户ID
	UserId int `mson:"UserId" json:"userid"`
	// 用户昵称
	Nickname string `mson:"NickName" json:"nickname"`
	// 头像地址
	Avatar string `mson:"Avatar" json:"avatar"`
	// 用户角色编码集合
	Roles []string `json:"roles"`
	// 用户权限标识集合
	Perms []string `json:"perms"`
}
