package entities

// Cycle OKR 周期
type Cycle struct {
	ID   int32  `gorm:"type:int not null;primaryKey"`
	Name string `gorm:"type:string not null"`
}
