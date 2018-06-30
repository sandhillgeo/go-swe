package swe

type Sensor struct {
	Id       int    `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"column:sensor_name;not null"`
	Comments string `gorm:"column:comments;not null"`
	Geometry []byte `gorm:"column:geom;not null"`
}

func (Sensor) TableName() string {
	return "sensor"
}
