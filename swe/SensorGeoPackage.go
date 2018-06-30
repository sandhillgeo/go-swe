package swe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

import (
	"github.com/sandhillgeo/go-gpkg/gpkg"
)

type SensorGeoPackage struct {
	*gpkg.GeoPackage
}

func NewSensorGeoPackage(uri string) *SensorGeoPackage {
	return &SensorGeoPackage{
		&gpkg.GeoPackage{
			Uri: uri,
		},
	}
}

func (sgpkg *SensorGeoPackage) AutoMigrate() {
	sgpkg.GeoPackage.AutoMigrate()
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedAirTemperature{})
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedAirTemperatureMapping{})
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedBlob{})
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedBlobMapping{})
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedWindSpeed{})
	sgpkg.GeoPackage.DB.AutoMigrate(ObservedWindSpeedMapping{})
	sgpkg.GeoPackage.DB.AutoMigrate(Sensor{})
	sgpkg.GeoPackage.DB.AutoMigrate(SensorConnectionInformation{})
	sgpkg.GeoPackage.DB.AutoMigrate(SensorConnectionInformationMapping{})
	sgpkg.GeoPackage.DB.AutoMigrate(SensorObservableProperty{})
	sgpkg.GeoPackage.DB.AutoMigrate(SensorObservablePropertyMapping{})
}

func (sgpkg *SensorGeoPackage) GetSensorsWithinDistance(longitude float64, latitude float64, distance float64) (*SensorList, error) {
	// TODO Iterate through all features and see if within certain distance

	sensors := make([]Sensor, 0)
	err := sgpkg.GeoPackage.DB.Find(&sensors).Error
	if err != nil {
		return &SensorList{}, err
	}
	return &SensorList{sensors: sensors}, err
}

func (sgpkg *SensorGeoPackage) GetConnectionInformation(s *Sensor) (*SensorConnectionInformation, error) {

	sensorConnectionInformationMapping := &SensorConnectionInformationMapping{}
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).First(sensorConnectionInformationMapping).Error
	if err != nil {
		return &SensorConnectionInformation{}, err
	}

	sensorConnectionInformation := &SensorConnectionInformation{}
	err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorConnectionInformationMapping.RelatedId}).First(sensorConnectionInformation).Error
	if err != nil {
		return &SensorConnectionInformation{}, err
	}

	return sensorConnectionInformation, nil
}

func (sgpkg *SensorGeoPackage) GetObservableProperties(s *Sensor) (*SensorObservablePropertyList, error) {

	sensorObservablePropertyMappings := make([]SensorObservablePropertyMapping, 0)
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).Find(&sensorObservablePropertyMappings).Error
	if err != nil {
		return &SensorObservablePropertyList{}, err
	}

	sensorObservableProperties := make([]SensorObservableProperty, len(sensorObservablePropertyMappings))
	for i, sensorObservablePropertyMapping := range sensorObservablePropertyMappings {
		sensorObservableProperty := SensorObservableProperty{}
		err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorObservablePropertyMapping.RelatedId}).First(&sensorObservableProperty).Error
		if err != nil {
			return &SensorObservablePropertyList{}, err
		}
		sensorObservableProperties[i] = sensorObservableProperty
	}

	return &SensorObservablePropertyList{sensorObservableProperties: sensorObservableProperties}, nil
}

func (spkg *SensorGeoPackage) UrlGetCapabilities(sci *SensorConnectionInformation) string {
	return sci.SensorObservationServiceEndpoint + "?service=SOS&version=" + sci.SensorObservationServiceVersion + "&request=GetCapabilities"
}

// TODO: UrlDescribeSensor

func (spkg *SensorGeoPackage) UrlGetResultTemplate(sci *SensorConnectionInformation, sop *SensorObservableProperty) string {
	return sci.SensorObservationServiceEndpoint + "?service=SOS&version=" + sci.SensorObservationServiceVersion + "&request=GetResultTemplate&offering=" + sop.SensorIdentifier + "&observableProperty=" + sop.SensorObservableProperty + "&responseFormat=application/json"
}

func (spkg *SensorGeoPackage) UrlGetResult(sci *SensorConnectionInformation, sop *SensorObservableProperty, rt *ResultTemplate) string {
	u := sci.SensorObservationServiceEndpoint + "?service=SOS&version=" + sci.SensorObservationServiceVersion + "&request=GetResult&offering=" + sop.SensorIdentifier + "&observableProperty=" + sop.SensorObservableProperty

	if sop.TimeEnd == "now" {
		u += "&temporalFilter=phenomenonTime,now"
	} else {
		u += "&temporalFilter=phenomenonTime," + sop.TimeStart + "/" + sop.TimeEnd
	}

	switch sop.GetDataType() {
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
		return body, err
	}

	return body, nil
}

func (spkg *SensorGeoPackage) RequestGetCapabilities(sci *SensorConnectionInformation) (*CapabilitiesDocument, error) {
	capabilitiesDocument := &CapabilitiesDocument{}

	url := spkg.UrlGetCapabilities(sci)

	body, err := spkg.MakeRequest(url)
	if err != nil {
		return capabilitiesDocument, err
	}

	err = json.Unmarshal(body, capabilitiesDocument)
	if err != nil {
		return capabilitiesDocument, err
	}

	return capabilitiesDocument, nil
}

func (sgpkg *SensorGeoPackage) RequestGetResultTemplate(sci *SensorConnectionInformation, sop *SensorObservableProperty) (*ResultTemplate, error) {
	resultTemplate := &ResultTemplate{}

	url := sgpkg.UrlGetResultTemplate(sci, sop)

	body, err := sgpkg.MakeRequest(url)
	if err != nil {
		return resultTemplate, err
	}

	err = json.Unmarshal(body, resultTemplate)
	if err != nil {
		return resultTemplate, err
	}

	return resultTemplate, nil
}

func (sgpkg *SensorGeoPackage) RequestGetResult(sci *SensorConnectionInformation, sop *SensorObservableProperty, rt *ResultTemplate) (interface{}, error) {
	var result interface{}

	url := sgpkg.UrlGetResult(sci, sop, rt)

	body, err := sgpkg.MakeRequest(url)
	if err != nil {
		return result, err
	}

	switch sop.SensorObservableProperty {
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

func (sgpkg *SensorGeoPackage) Run() error {

	sensors, err := sgpkg.GetSensorsWithinDistance(0, 0, 0)
	if err != nil {
		return err
	}

	for i := 0; i < sensors.Size(); i++ {
		sensor := sensors.Item(i)
		sci, err := sgpkg.GetConnectionInformation(sensor)
		if err != nil {
			return err
		}

		_, err = sgpkg.RequestGetCapabilities(sci)
		if err != nil {
			return err
		}

		observableProperties, err := sgpkg.GetObservableProperties(sensor)
		if err != nil {
			return err
		}

		for j := 0; j < observableProperties.Size(); j++ {
			observableProperty := observableProperties.Item(j)

			resultTemplate, err := sgpkg.RequestGetResultTemplate(sci, observableProperty)
			if err != nil {
				return err
			}

			result, err := sgpkg.RequestGetResult(sci, observableProperty, resultTemplate)
			if err != nil {
				return err
			}

			switch result.(type) {
			case ResultAirTemperature:
				result2 := result.(ResultAirTemperature)
				sgpkg.GeoPackage.DB.Create(&ObservedAirTemperature{Time: result2.Time, Temperature: result2.Temperature})
			case ResultWindSpeed:
				result2 := result.(ResultWindSpeed)
				sgpkg.GeoPackage.DB.Create(&ObservedWindSpeed{Time: result2.Time, WindSpeed: result2.WindSpeed})
			case ResultBlob:
				result2 := result.(ResultBlob)
				sgpkg.GeoPackage.DB.Create(&ObservedBlob{Blob: result2.Blob})
			}
		}

	}

	return nil

}
