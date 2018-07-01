package swe

type SensorMapByWifiNetwork struct {
	sensors map[WifiNetwork]SensorList
}

func (m *SensorMapByWifiNetwork) Keys() *WifiNetworkList {
	networks := make([]WifiNetwork, 0, len(m.sensors))
	for network := range m.sensors {
		networks = append(networks, network)
	}
	return &WifiNetworkList{networks: networks}
}

func (m *SensorMapByWifiNetwork) Get(key *WifiNetwork) *SensorList {
	sensorList := m.sensors[*key]
	return &sensorList
}
