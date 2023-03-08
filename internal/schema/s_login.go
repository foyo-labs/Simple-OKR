package schema

type LoginParam struct {
	Email    string `json:"email" binding:"required"`    // 用户名
	Password string `json:"password" binding:"required"` // 密码(md5加密)
}

type LoginTokenInfo struct {
	AccessToken string   `json:"access_token"` // 访问令牌
	UserID      string   `json:"user_id"`      // 用户编号
	UserInfo    UserInfo `json:"user_info"`    // 用户信息
}

// UserInfo 用户信息，根据需要扩展必要字段
type UserInfo struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
}
