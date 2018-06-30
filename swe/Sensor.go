package swe

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

import (
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/wkb"
)
import (
  "github.com/sandhillgeo/go-gpkg/gpkg"
)


type Sensor struct {
  Id int        `gorm:"column:id;not null;primary_key"`
  Name string `gorm:"column:sensor_name;not null"`
  Comments string `gorm:"column:comments;not null"`
  Geometry []byte `gorm:"column:geom;not null"`
}

func (Sensor) TableName() string {
	return "sensor"
}
