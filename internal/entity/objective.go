package entity

type Objective struct {
	ID          int32  `gorm:"column:id;not null;primaryKey"`
	Name        string `gorm:"column:name;not null"`
	Description string `gorm:"column:discription"`
	Actived     bool   `gorm:"column:actived;not null"`
	Sequence    int32  `gorm:"column:sequence"`
	ParentID    int32  `gorm:"column:parent_id"`
}
