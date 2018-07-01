package swe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

import (
	"github.com/pkg/errors"
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

func (sgpkg *SensorGeoPackage) AutoMigrate() error {
	err := sgpkg.GeoPackage.AutoMigrate()
	if err != nil {
		return errors.Wrap(err, "Error migrating GeoPackage tables")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedAirTemperature{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedAirTemperature")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedAirTemperatureMapping{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedAirTemperatureMapping")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedBlob{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedBlob")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedBlobMapping{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedBlobMapping")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedWindSpeed{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedWindSpeed")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(ObservedWindSpeedMapping{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating ObservedWindSpeedMapping")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(Sensor{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating Sensor")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(SensorConnectionInformation{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating SensorConnectionInformation")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(SensorConnectionInformationMapping{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating SensorConnectionInformationMapping")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(SensorObservableProperty{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating SensorObservableProperty")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(SensorObservablePropertyMapping{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating SensorObservablePropertyMapping")
	}
	err = sgpkg.GeoPackage.DB.AutoMigrate(gpkg.Relation{}).Error
	if err != nil {
		return errors.Wrap(err, "Error migrating Relation")
	}

	relation := gpkg.Relation{
		BaseTableName:        Sensor{}.TableName(),
		BasePrimaryColumn:    "id",
		RelatedTableName:     SensorConnectionInformation{}.TableName(),
		RelatedPrimaryColumn: "id",
		RelationName:         "simple_attributes",
		MappingTableName:     SensorConnectionInformationMapping{}.TableName(),
	}
	err = sgpkg.GeoPackage.DB.Create(&relation).Error
	if err != nil {
		return errors.Wrap(err, "Error creating relation "+fmt.Sprint(relation))
	}

	relation = gpkg.Relation{
		BaseTableName:        Sensor{}.TableName(),
		BasePrimaryColumn:    "id",
		RelatedTableName:     SensorObservableProperty{}.TableName(),
		RelatedPrimaryColumn: "id",
		RelationName:         "simple_attributes",
		MappingTableName:     SensorObservablePropertyMapping{}.TableName(),
	}
	err = sgpkg.GeoPackage.DB.Create(&relation).Error
	if err != nil {
		return errors.Wrap(err, "Error creating relation "+fmt.Sprint(relation))
	}

	relation = gpkg.Relation{
		BaseTableName:        SensorObservableProperty{}.TableName(),
		BasePrimaryColumn:    "id",
		RelatedTableName:     ObservedAirTemperature{}.TableName(),
		RelatedPrimaryColumn: "id",
		RelationName:         "simple_attributes",
		MappingTableName:     ObservedAirTemperatureMapping{}.TableName(),
	}
	err = sgpkg.GeoPackage.DB.Create(&relation).Error
	if err != nil {
		return errors.Wrap(err, "Error creating relation "+fmt.Sprint(relation))
	}

	relation = gpkg.Relation{
		BaseTableName:        SensorObservableProperty{}.TableName(),
		BasePrimaryColumn:    "id",
		RelatedTableName:     ObservedWindSpeed{}.TableName(),
		RelatedPrimaryColumn: "id",
		RelationName:         "simple_attributes",
		MappingTableName:     ObservedWindSpeedMapping{}.TableName(),
	}
	err = sgpkg.GeoPackage.DB.Create(&relation).Error
	if err != nil {
		return errors.Wrap(err, "Error creating relation "+fmt.Sprint(relation))
	}

	relation = gpkg.Relation{
		BaseTableName:        SensorObservableProperty{}.TableName(),
		BasePrimaryColumn:    "id",
		RelatedTableName:     ObservedBlob{}.TableName(),
		RelatedPrimaryColumn: "id",
		RelationName:         "media",
		MappingTableName:     ObservedBlobMapping{}.TableName(),
	}
	err = sgpkg.GeoPackage.DB.Create(&relation).Error
	if err != nil {
		return errors.Wrap(err, "Error creating relation "+fmt.Sprint(relation))
	}

	table_names := []string{
		gpkg.Relation{}.TableName(),
		ObservedAirTemperatureMapping{}.TableName(),
		ObservedWindSpeedMapping{}.TableName(),
		ObservedBlobMapping{}.TableName(),
		SensorConnectionInformationMapping{}.TableName(),
		SensorObservablePropertyMapping{}.TableName(),
	}

	sgpkg.GeoPackage.DB.AutoMigrate(gpkg.Extension{})
	for _, table_name := range table_names {
		extension := gpkg.Extension{
			Table:      table_name,
			Column:     nil,
			Extension:  "related_tables",
			Definition: "TBD",
			Scope:      "read-write",
		}
		err = sgpkg.GeoPackage.DB.Create(&extension).Error
		if err != nil {
			return errors.Wrap(err, "Error creating extension "+fmt.Sprint(extension))
		}
	}
	return nil
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

func (sgpkg *SensorGeoPackage) GetSensorsWithinDistanceByWifiNetwork(longitude float64, latitude float64, distance float64) (*SensorMapByWifiNetwork, error) {

	sensorMapByWifiNetwork := &SensorMapByWifiNetwork{}

	sensorList, err := sgpkg.GetSensorsWithinDistance(longitude, latitude, distance)
	if err != nil {
		return sensorMapByWifiNetwork, err
	}

	m := map[*WifiNetwork][]Sensor{}
	for i := 0; i < sensorList.Size(); i++ {
		sensor := sensorList.Item(i)
		sensorConnectionInformation, err := sgpkg.GetConnectionInformation(sensor)
		if err != nil {
			return sensorMapByWifiNetwork, err
		}
		var network *WifiNetwork
		for k, _ := range m {
			if sensorConnectionInformation.WiFiSSID == k.WiFiSSID && sensorConnectionInformation.WiFiPassword == k.WiFiPassword {
				network = k
			}
		}
		if network == nil {
			network = &WifiNetwork{WiFiSSID: sensorConnectionInformation.WiFiSSID, WiFiPassword: sensorConnectionInformation.WiFiPassword}
			m[network] = make([]Sensor, 0)
		}
		m[network] = append(m[network], *sensor)
	}

	m2 := map[WifiNetwork]SensorList{}
	for network, sensors := range m {
		m2[*network] = SensorList{sensors: sensors}
	}

	return &SensorMapByWifiNetwork{sensors: m2}, err
}

func (sgpkg *SensorGeoPackage) GetConnectionInformation(s *Sensor) (*SensorConnectionInformation, error) {

	sensorConnectionInformationMapping := &SensorConnectionInformationMapping{}
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).First(sensorConnectionInformationMapping).Error
	if err != nil {
		return &SensorConnectionInformation{}, errors.Wrap(err, "Error selecting SensorConnectionInformationMapping with base_id "+fmt.Sprint(s.Id))
	}

	sensorConnectionInformation := &SensorConnectionInformation{}
	err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorConnectionInformationMapping.RelatedId}).First(sensorConnectionInformation).Error
	if err != nil {
		return &SensorConnectionInformation{}, errors.Wrap(err, "Error selecting SensorConnectionInformation with id "+fmt.Sprint(sensorConnectionInformationMapping.RelatedId))
	}

	return sensorConnectionInformation, nil
}

func (sgpkg *SensorGeoPackage) GetObservableProperties(s *Sensor) (*SensorObservablePropertyList, error) {

	sensorObservablePropertyMappings := make([]SensorObservablePropertyMapping, 0)
	err := sgpkg.GeoPackage.DB.Where(map[string]interface{}{"base_id": s.Id}).Find(&sensorObservablePropertyMappings).Error
	if err != nil {
		return &SensorObservablePropertyList{}, errors.Wrap(err, "Error selecting SensorObservablePropertyMapping with base_id "+fmt.Sprint(s.Id))
	}

	sensorObservableProperties := make([]SensorObservableProperty, len(sensorObservablePropertyMappings))
	for i, sensorObservablePropertyMapping := range sensorObservablePropertyMappings {
		sensorObservableProperty := SensorObservableProperty{}
		err = sgpkg.GeoPackage.DB.Where(map[string]interface{}{"id": sensorObservablePropertyMapping.RelatedId}).First(&sensorObservableProperty).Error
		if err != nil {
			return &SensorObservablePropertyList{}, errors.Wrap(err, "Error selecting SensorObservableProperty with id "+fmt.Sprint(sensorObservablePropertyMapping.RelatedId))
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

	if sop.TimeStart == "" && sop.TimeEnd == "now" {
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

func (sgpkg *SensorGeoPackage) Sync(network *WifiNetwork, sensors *SensorList) error {

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

			sgpkg.AddObservation(observableProperty, result)
		}
	}

	return nil

}

func (sgpkg *SensorGeoPackage) AddObservation(observableProperty *SensorObservableProperty, result interface{}) error {
	switch result.(type) {
	case ResultAirTemperature:
		result2 := result.(ResultAirTemperature)
		observedAirTemperature := ObservedAirTemperature{Time: result2.Time, Temperature: result2.Temperature}
		err := sgpkg.GeoPackage.DB.Create(&observedAirTemperature).Error
		if err != nil {
			return err
		}
		sgpkg.GeoPackage.DB.Create(&ObservedAirTemperatureMapping{BaseId: observableProperty.Id, RelatedId: observedAirTemperature.Id})
	case ResultWindSpeed:
		result2 := result.(ResultWindSpeed)
		observedWindSpeed := ObservedWindSpeed{Time: result2.Time, WindSpeed: result2.WindSpeed}
		err := sgpkg.GeoPackage.DB.Create(&observedWindSpeed).Error
		if err != nil {
			return err
		}
		sgpkg.GeoPackage.DB.Create(&ObservedWindSpeedMapping{BaseId: observableProperty.Id, RelatedId: observedWindSpeed.Id})
	case ResultBlob:
		result_blob := result.(ResultBlob)
		observedBlob := ObservedBlob{Blob: result_blob.Blob}
		err := sgpkg.GeoPackage.DB.Create(&observedBlob).Error
		if err != nil {
			return err
		}
		sgpkg.GeoPackage.DB.Create(&ObservedBlobMapping{BaseId: observableProperty.Id, RelatedId: observedBlob.Id})
	}
	return nil
}
