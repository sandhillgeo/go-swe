package swe

type SensorConnectionInformationMapping struct {
	Id        int `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	BaseId    int `gorm:"column:base_id;not null"`
	RelatedId int `gorm:"column:related_id;not null"`
}

func (SensorConnectionInformationMapping) TableName() string {
	return "sensor_connection_information_mapping"
}
