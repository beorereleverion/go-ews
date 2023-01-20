package elements

// The Periods element represents an array of periods that define the time offset at different stages of the time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/periods
import "encoding/xml"

type Periods struct {
	XMLName xml.Name

	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
}

func (P *Periods) SetForMarshal() {
	P.XMLName.Local = "t:Periods"
}

func (P *Periods) GetSchema() *Schema {
	return &SchemaTypes
}
