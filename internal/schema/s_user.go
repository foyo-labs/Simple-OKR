package schema

type User struct {
	ID           string `json:"id"`            // 唯一标识
	Email        string `json:"user_name"`     // 邮箱
	Name         string `json:"real_name"`     // 姓名
	Password     string `json:"password"`      // 密码
	Status       int    `json:"status"`        // 状态
	Role         Role   `json:"role"`          // 角色
	DepartmentID string `json:"department_id"` //部门编号
	Created      uint64 `json:"created"`       // 创建时间
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

type ReqestRegistion struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	DepartmentID string `json:"department_id"` //用于某个用户开通当前企业用户，只有角色为Admin的有权限
}
