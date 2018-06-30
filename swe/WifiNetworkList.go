package swe

type WifiNetworkList struct {
	networks []WifiNetwork
}

func (l *WifiNetworkList) Size() int {
	return len(l.networks)
}

func (l *WifiNetworkList) Item(i int) *WifiNetwork {
	return &l.networks[i]
}
