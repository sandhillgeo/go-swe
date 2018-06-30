package swe

type ObservedBlobMapping struct {
  BaseId int `gorm:"column:base_id;not null"`
  RelatedId int `gorm:"column:related_id;not null"`
}

func (ObservedBlobMapping) TableName() string {
	return "observed_blob_mapping"
}
