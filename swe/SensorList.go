package swe

type SensorList struct {
	sensors []Sensor
}

func (l *SensorList) Size() int {
	return len(l.sensors)
}

func (l *SensorList) Item(i int) *Sensor {
	return &l.sensors[i]
}
