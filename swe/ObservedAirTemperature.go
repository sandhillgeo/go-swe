package swe

import (
	"time"
)

type ObservedAirTemperature struct {
	Id          int       `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	Time        time.Time `json:"time"`
	Temperature float64   `json:"temperature"`
}

func (ObservedAirTemperature) TableName() string {
	return "observed_air_temperature"
}
