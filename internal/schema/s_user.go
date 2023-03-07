package schema

type User struct {
	ID       string `json:"id"`        // 唯一标识
	Email    string `json:"user_name"` // 邮箱
	Name     string `json:"real_name"` // 姓名
	Password string `json:"password"`  // 密码
	Status   int    `json:"status"`
	Created  uint64 `json:"created"` // 创建时间
}

type Users []*User

type UserQueryParam struct {
	PaginationParam
	Email  string
	Status int
}

type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}
