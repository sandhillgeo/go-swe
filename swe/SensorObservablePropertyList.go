package swe

type SensorObservablePropertyList struct {
	sensorObservableProperties []SensorObservableProperty
}

func (l *SensorObservablePropertyList) Size() int {
	return len(l.sensorObservableProperties)
}

func (l *SensorObservablePropertyList) Item(i int) *SensorObservableProperty {
	return &l.sensorObservableProperties[i]
}
