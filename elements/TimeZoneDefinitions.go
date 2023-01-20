package elements

// The TimeZoneDefinitions element represents an array of time zone definitions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezonedefinitions
import "encoding/xml"

type TimeZoneDefinitions struct {
	XMLName xml.Name

	// The TimeZoneDefinition element specifies the periods and transitions that define a time zone.
	TimeZoneDefinition *TimeZoneDefinition `xml:"TimeZoneDefinition"`
}

func (T *TimeZoneDefinitions) SetForMarshal() {
	T.XMLName.Local = "m:TimeZoneDefinitions"
}

func (T *TimeZoneDefinitions) GetSchema() *Schema {
	return &SchemaMessages
}
