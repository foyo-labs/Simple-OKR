package entities

type Department struct {
	ID       int32  `gorm:"type:int not null;primaryKey"`
	Name     string `gorm:"type:string not null"`
	ParentID int32  `gorm:"type:int not null"`
}

type Company struct {
	ID   int32  `gorm:"type:int not null;primaryKey"`
	Name string `gorm:"type:string not null"`
	Abbr string `gorm:"type:string"`
}
