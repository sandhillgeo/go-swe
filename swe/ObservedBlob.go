package swe

type ObservedBlob struct {
	Id   int    `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	Blob []byte `json:"blob"`
}

func (ObservedBlob) TableName() string {
	return "observed_blob"
}
