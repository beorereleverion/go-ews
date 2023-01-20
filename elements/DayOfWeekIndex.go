package elements

// The DayOfWeekIndex element describes which week in a month is used in a relative recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayofweekindex
import "encoding/xml"

type DayOfWeekIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DayOfWeekIndex) SetForMarshal() {
	D.XMLName.Local = "t:DayOfWeekIndex"
}

func (D *DayOfWeekIndex) GetSchema() *Schema {
	return &SchemaTypes
}
