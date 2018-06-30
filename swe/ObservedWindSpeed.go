package swe

type ObservedWindSpeed struct {
  Id int        `gorm:"column:id;not null;primary_key"`
  Time time.Time `json:"time"`
  WindSpeed float64 `json:"wind_speed"`
}

func (ObservedWindSpeed) TableName() string {
	return "observed_wind_speed"
}
