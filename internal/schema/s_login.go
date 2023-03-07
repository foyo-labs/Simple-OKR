package schema

type LoginParam struct {
	Email    string `json:"email" binding:"required"`    // 用户名
	Password string `json:"password" binding:"required"` // 密码(md5加密)
}

type LoginTokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	UserID      string `json:"user_id"`      // 用户编号
}
