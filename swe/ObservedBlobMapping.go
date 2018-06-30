package swe

type ObservedBlobMapping struct {
	Id        int `gorm:"column:id;not null;unique;primary_key;AUTO_INCREMENT"`
	BaseId    int `gorm:"column:base_id;not null"`
	RelatedId int `gorm:"column:related_id;not null"`
}

func (ObservedBlobMapping) TableName() string {
	return "observed_blob_mapping"
}
