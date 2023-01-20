package elements

// The DayOfMonth element describes the day in a month that a recurring item occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayofmonth
import "encoding/xml"

type DayOfMonth struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (D *DayOfMonth) SetForMarshal() {
	D.XMLName.Local = "t:DayOfMonth"
}

func (D *DayOfMonth) GetSchema() *Schema {
	return &SchemaTypes
}
