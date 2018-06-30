package swe

type SensorObservablePropertyMapping struct {
  BaseId int `gorm:"column:base_id;not null"`
  RelatedId int `gorm:"column:related_id;not null"`
}

func (SensorObservableProperty) TableName() string {
	return "sensor_observable_property_mapping"
}
