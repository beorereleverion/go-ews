package elements

// The Properties element contains a list of one or more tracking properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/properties-arrayoftrackingpropertiestype
import "encoding/xml"

type PropertiesArrayOfTrackingPropertiesType struct {
	XMLName xml.Name

	// The TrackingPropertyType element represents a name and value pair of strings that is used to create properties for message tracking reports.
	TrackingPropertyType *TrackingPropertyType `xml:"TrackingPropertyType"`
}

func (P *PropertiesArrayOfTrackingPropertiesType) SetForMarshal() {
	P.XMLName.Local = "m:Properties"
}

func (P *PropertiesArrayOfTrackingPropertiesType) GetSchema() *Schema {
	return &SchemaMessages
}
