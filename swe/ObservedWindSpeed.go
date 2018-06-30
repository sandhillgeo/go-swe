package swe

import (
	"time"
)

type ObservedWindSpeed struct {
	Id        int       `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	Time      time.Time `json:"time"`
	WindSpeed float64   `json:"wind_speed"`
}

func (ObservedWindSpeed) TableName() string {
	return "observed_wind_speed"
}
