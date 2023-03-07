package entity

type KeyResult struct {
	ID           int32   `gorm:"column:id;not null;primaryKey"`
	Name         string  `gorm:"column:name;not null"`
	StartValue   float64 `gorm:"column:start_value;not null;default:0;"`
	TargetValue  float64 `gorm:"column:target_value;not null;default:0;"`
	CurrentValue float64 `gorm:"column:current_value;not null;default:0;"`
	Sequence     int32   `gorm:"column:sequence;not null"`
	Created      uint64  `gorm:"column:created;index;"`
}
