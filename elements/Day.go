package elements

// The Day element represents the day of the month on which the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/day
import "encoding/xml"

type Day struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (D *Day) SetForMarshal() {
	D.XMLName.Local = "t:Day"
}

func (D *Day) GetSchema() *Schema {
	return &SchemaTypes
}
