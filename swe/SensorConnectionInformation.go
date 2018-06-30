package swe

type SensorConnectionInformation struct {
	Id                               int    `gorm:"column:id;not null;primary_key"`
	SensorIdentifier                 string `gorm:"column:sensor_identifier;not null"`
	SensorObservationServiceEndpoint string `gorm:"column:sos_endpoint;not null"`
	SensorObservationServiceVersion  string `gorm:"column:sos_version;not null"`
	WiFiSSID                         string `gorm:"column:wifi_ssid;not null"`
	WiFiPassword                     string `gorm:"column:wifi_password"`
}

func (SensorConnectionInformation) TableName() string {
	return "sensor_connection_information"
}
