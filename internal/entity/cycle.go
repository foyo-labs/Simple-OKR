package entity

// Cycle OKR 周期
type Cycle struct {
	ID   int32  `gorm:"column:id;not null;primaryKey"`
	Name string `gorm:"column:name;not null"`
}
