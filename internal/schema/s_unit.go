package schema

type CreateUserDepartment struct {
	CompanyID string `json:"company_id"`
}

type Company struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Abbr    string `json:"abbr"`
	Created uint64 `json:"created"`
}
