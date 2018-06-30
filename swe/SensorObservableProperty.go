package swe

type SensorObservableProperty struct {
	Id                       int    `gorm:"column:id;not null;primary_key"`
	SensorIdentifier         string `gorm:"column:sensor_identifier;not null"`
	SensorProcedure          string `gorm:"column:sensor_procedure;not null"`
	SensorObservableProperty string `gorm:"column:observable_property;not null"`
	TimeStart                string `gorm:"column:time_start;not null"`
	TimeEnd                  string `gorm:"column:time_end;not null"`
}

func (sop *SensorObservableProperty) GetDataType() string {
	switch sop.SensorObservableProperty {
	case "http://sensorml.com/ont/swe/property/AirTemperature":
		return "float"
	case "http://sensorml.com/ont/swe/property/WindSpeed":
		return "float"
	case "http://sensorml.com/ont/swe/property/Location":
		return "location"
	case "http://www.opengis.net/def/property/OGC/0/SensorLocation":
		return "location"
	case "http://sensorml.com/ont/swe/property/VideoFrame":
		return "blob"
	}
	return ""
}

func (SensorObservableProperty) TableName() string {
	return "sensor_observable_property"
}
