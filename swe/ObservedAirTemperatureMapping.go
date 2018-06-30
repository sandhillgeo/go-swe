package swe

type ObservedAirTemperatureMapping struct {
	Id        int `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	BaseId    int `gorm:"column:base_id;not null"`
	RelatedId int `gorm:"column:related_id;not null"`
}

func (ObservedAirTemperatureMapping) TableName() string {
	return "observed_air_temperature_mapping"
}
