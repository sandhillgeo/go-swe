package swe

type ObservedAirTemperatureMapping struct {
	BaseId    int `gorm:"column:base_id;not null"`
	RelatedId int `gorm:"column:related_id;not null"`
}

func (ObservedAirTemperatureMapping) TableName() string {
	return "observed_air_temperature_mapping"
}
