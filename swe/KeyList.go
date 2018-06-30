package swe

type KeyList struct {
	keys []string
}

func (l *KeyList) Size() int {
	return len(l.keys)
}

func (l *KeyList) Item(i int) string {
	return l.keys[i]
}
