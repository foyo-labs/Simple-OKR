package entities

type User struct {
	ID       int32  `gorm:"type:int not null;primaryKey"`
	Email    string `gorm:"type:string not null"`
	Password string `gorm:"type:string not null"`
	Name     string `gorm:"type:string"`
	Created  uint64 `gorm:"type:int not null"`
}

type AdminUser struct {
	ID      int32  `gorm:"type:int not null;primaryKey"`
	Created uint64 `gorm:"type:int not null"`
}

type AuditorUser struct {
	ID int32 `gorm:"type:int not null;primaryKey"`
}

type UserSettings struct {
	ID           int32 `gorm:"type:int not null;primaryKey"`
	UserID       int32 `gorm:"type:int not null"`
	CompanyID    int32 `gorm:"type:int not null"`
	DepartmentID int32 `gorm:"type:int not null"`
}
