package elements

// The IsAllDayEvent element indicates whether a calendar item or meeting request represents an all-day event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isalldayevent
import "encoding/xml"

type IsAllDayEvent struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsAllDayEvent) SetForMarshal() {
	I.XMLName.Local = "t:IsAllDayEvent"
}

func (I *IsAllDayEvent) GetSchema() *Schema {
	return &SchemaTypes
}
