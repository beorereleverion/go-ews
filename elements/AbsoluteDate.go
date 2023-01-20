package elements

// The AbsoluteDate element represents the date when the time changes from standard or daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absolutedate
import "time"

import "encoding/xml"

type AbsoluteDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (A *AbsoluteDate) SetForMarshal() {
	A.XMLName.Local = "t:AbsoluteDate"
}

func (A *AbsoluteDate) GetSchema() *Schema {
	return &SchemaTypes
}
