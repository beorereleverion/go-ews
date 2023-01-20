package elements

// The Month element represents the month in which the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/month-time-zone-transition
import "encoding/xml"

type MonthTimeZoneTransition struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MonthTimeZoneTransition) SetForMarshal() {
	M.XMLName.Local = "t:Month"
}

func (M *MonthTimeZoneTransition) GetSchema() *Schema {
	return &SchemaTypes
}
