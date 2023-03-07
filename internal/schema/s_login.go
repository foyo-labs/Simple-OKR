package schema

type LoginParam struct {
	Email    string `json:"email" binding:"required"`    // 用户名
	Password string `json:"password" binding:"required"` // 密码(md5加密)
}

type LoginTokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 令牌到期时间戳
	UserID      string `json:"user_id"`      // 用户编号
	Email       string `json:"email"`        // 邮箱
}
