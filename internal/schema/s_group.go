package schema

type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ParentID     string `json:"parent_id"`
	LevelNum     string `json:"level_num"` //分级设置：LIKE　0.1.3
	OrderNum     int64  `json:"order_num"` //排序
	CurrentLevel uint64 `json:"curr_lev"`  //当前级别,第一级为0，依次增长
	Created      uint64 `json:"created"`
	UserID       string `json:"user_id"`
}

type UserGroup struct {
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`
}

type CreateGroupRequest struct {
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
}
