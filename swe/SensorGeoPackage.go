package swe

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

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


type SensorGeoPackage struct {
  *gpkg.GeoPackage
}

func New(uri string) *SensorGeoPackage {
	return &SensorGeoPackage{
    &GeoPackage{
  		Uri: uri,
  	},
  }
}

func (sgpkg *SensorGeoPackage) GetSensorsWithinDistance(longitude float64, latitude float64, distance float64) (*SensorList, err) {
  // TODO Iterate through all features and see if within certain distance

	sensors := make([]Sensor, 0)
	err := sgpkg.GeoPackage.DB.Find(&sensors)
	if err != nil {
		return &SensorList{}, err
	}
	return &SensorList{sensors: sensors}, err
}


func (sgpkg *SensorGeoPackage) GetConnectionInformation(s *Sensor) (*SensorConnectionInformation, error) {

  sensorConnectionInformationMapping := &SensorConnectionInformationMapping{}
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).First(sensorConnectionInformationMapping)
	if err != nil {
		return err
	}

	sensorConnectionInformation = &SensorConnectionInformation{}
	err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorConnectionInformationMapping.RelatedId}).First(sensorConnectionInformation)
	if err != nil {
		return err
	}

	return sensorConnectionInformation, nil
}

func (s *Sensor) GetObservableProperties(s *Sensor) (*SensorObservablePropertyList, error) {

	sensorObservablePropertyMappings := make([]SensorObservablePropertyMapping, 0)
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).Find(&sensorObservablePropertyMappings)
	if err != nil {
		return err
	}

	sensorObservableProperties = make([]SensorObservableProperty, len(sensorObservablePropertyMappings))
	for i, sensorObservablePropertyMapping := range {
		sensorObservableProperty := SensorObservableProperty{}
		err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorObservablePropertyMapping.RelatedId}).First(&sensorObservableProperty)
		if err != nil {
			return err
		}
		sensorObservableProperties[i] = sensorObservableProperty
	}

	return SensorObservablePropertyList{sensorObservableProperties: sensorObservableProperties}, nil
}

func (spkg *SensorGeoPackage) UrlGetCapabilities(sci *SensorConnectionInformation) string {
  return sci.SensorObservationServiceEndpoint+"?service=SOS&version="+SensorObservationServiceVersion+"&request=GetCapabilities"
}

// TODO: UrlDescribeSensor

func (spkg *SensorGeoPackage) UrlGetResultTemplate(sci *SensorConnectionInformation, sop *SensorObservableProperty) string {
  return sci.SensorObservationServiceEndpoint+"?service=SOS&version="+SensorObservationServiceVersion+"&request=GetResultTemplate&offering="+sop.SensorIdentifier+"&observableProperty="+sop.SensorObservableProperty+"&responseFormat=application/json"
}

func (spkg *SensorGeoPackage) UrlGetResult(sci *SensorConnectionInformation, sop *SensorObservableProperty, rt *ResultTemplate) string {
	u := sci.SensorObservationServiceEndpoint+"?service=SOS&version="+SensorObservationServiceVersion+"&request=GetResult&offering="+sop.SensorIdentifier+"&observableProperty="+sop.SensorObservableProperty

	if sop.TimeEnd == "now" {
		u += "&temporalFilter=phenomenonTime,now"
	} else {
    u += "&temporalFilter=phenomenonTime,"+sop.TimeStart+"/"+sop.TimeEnd
	}

	switch (sop.GetDataType()) {
	case "float":
		u += "&responseFormat=application/json"
	case "location":
		u += "&responseFormat=application/json"
	}

	return u
}

func (spkg *SensorGeoPackage) MakeRequest(url string) ([]byte, error) {

	client := http.Client{
		Timeout: time.Second * 120,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	response, err := client.Do(request)
	if err != nil {
		return make([]byte, 0), err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return body, nil
}

func (spkg *SensorGeoPackage) RequestGetCapabilities(sci *SensorConnectionInformation) (*CapabilitiesDocument, error) {
	capabilitiesDocument = &CapabilitiesDocument{}

	url := spkg.UrlGetCapabilities(sci)

	body, err := MakeRequest(url)
	if err != nil {
		return capabilitiesDocument, error
	}

  err := json.Unmarshal(body, capabilitiesDocument)
	if err != nil {
		return capabilitiesDocument, error
	}

	return capabilitiesDocument, nil
}

func (spkg *SensorGeoPackage) RequestGetResultTemplate(sci *SensorConnectionInformation, sop *SensorObservableProperty) (*ResultTemplate, error) {
	resultTemplate = &ResultTemplate{}

	url := spkg.UrlGetResultTemplate(sci, sop)

	body, err := MakeRequest(url)
	if err != nil {
		return capabilitiesDocument, error
	}

  err := json.Unmarshal(body, resultTemplate)
	if err != nil {
		return resultTemplate, error
	}

	return resultTemplate, nil
}

func (spkg *SensorGeoPackage) RequestGetResult(sci *SensorConnectionInformation, sop *SensorObservableProperty, rt *ResultTemplate) (interface{}, error) {
	var result interface{}

	url := spkg.UrlGetResult(sci, sop, rt)

	body, err := MakeRequest(url)
	if err != nil {
		return result, error
	}

	switch (sop.SensorObservableProperty) {
  case "http://sensorml.com/ont/swe/property/AirTemperature":
		result := &ResultAirTemperature{}
		err := json.Unmarshal(body, result)
		return result, err
  case "http://sensorml.com/ont/swe/property/WindSpeed":
		result := &ResultWindSpeed{}
		err := json.Unmarshal(body, result)
		return result, err
  case "http://sensorml.com/ont/swe/property/Location":
		result := &ResultLocation{}
		err := json.Unmarshal(body, result)
		return result, err
  case "http://www.opengis.net/def/property/OGC/0/SensorLocation":
		result := &ResultLocation{}
		err := json.Unmarshal(body, result)
		return result, err
  case "http://sensorml.com/ont/swe/property/VideoFrame":
    return &ResultBlob{Blob: body}, nil
  }

	return &ResultBlob{Blob: body}, nil
}

func (spkg *SensorGeoPackage) Run() error {

  sensors, err := GetSensorsWithinDistance(0,0,0)
	if err != nil {
		return err
	}

  for _, sensor := range sensors {
		sci, err := spkg.GetConnectionInformation(sensor)
		if err != nil {
			return err
		}

		_, err := spkg.RequestGetCapabilities(sci)
		if err != nil {
			return err
		}

		observableProperties, err := spkg.GetObservableProperties(s)
		if err != nil {
			return err
		}

		for _, observableProperty := range observableProperties {

		    resultTemplate, err := spkg.RequestGetResultTemplate(sci, observableProperty)
				if err != nil {
					return err
				}

				result, err := spkg.RequestGetResult(sci, observableProperty, resultTemplate)
				if err != nil {
					return err
				}

				switch result.(type) {
				case ResultAirTemperature:
					result2 := result.(ResultAirTemperature)
					spkg.GeoPackage.DB.Create(&ObservedAirTemperature{Time: result2.Time, Temperature: result2.Temperature})
				case ResultWindSpeed:
					result2 := result.(ResultWindSpeed)
					spkg.GeoPackage.DB.Create(&ObservedWindSpeed{Time: result2.Time, WindSpeed: result2.WindSpeed})
				case ResultBlob:
					result2 := result.(ResultBlob)
					spkg.GeoPackage.DB.Create(&ObservedBlob{Blob: result2.Blob})
				}
		}

	}

	return nil

}
