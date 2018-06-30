package swe

import (
  "time"
)

type ResultAirTemperature struct {
	Time        time.Time `json:"time"`
	Temperature float64   `json:"temperature"`
}

type ResultWindSpeed struct {
	Time      time.Time `json:"time"`
	WindSpeed float64   `json:"windSpeed"`
}

type ResultBlob struct {
	Blob []byte
}

type ResultLocation struct {
	Time     time.Time               `json:"time"`
	Location *ResultLocationLocation `json:"location"`
}

type ResultLocationLocation struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
	Altitude  float64 `json:"alt"`
}
