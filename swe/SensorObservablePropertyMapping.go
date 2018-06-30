package swe

type SensorObservablePropertyMapping struct {
	Id        int `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	BaseId    int `gorm:"column:base_id;not null"`
	RelatedId int `gorm:"column:related_id;not null"`
}

func (SensorObservablePropertyMapping) TableName() string {
	return "sensor_observable_property_mapping"
}
