package swe

type Sensor struct {
	Id       int    `gorm:"column:id;not null;primary_key"`
	Name     string `gorm:"column:sensor_name;not null"`
	Comments string `gorm:"column:comments;not null"`
	Geometry []byte `gorm:"column:geom;not null"`
}

func (Sensor) TableName() string {
	return "sensor"
}
