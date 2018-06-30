package swe

type SensorConnectionInformationMapping struct {
  BaseId int `gorm:"column:base_id;not null"`
  RelatedId int `gorm:"column:related_id;not null"`
}

func (SensorConnectionInformation) TableName() string {
	return "sensor_connection_information_mapping"
}
