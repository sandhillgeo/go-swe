package swe

type ObservedWindSpeedMapping struct {
  BaseId int `gorm:"column:base_id;not null"`
  RelatedId int `gorm:"column:related_id;not null"`
}

func (ObservedWindSpeedMapping) TableName() string {
	return "wind_speed_mapping"
}
