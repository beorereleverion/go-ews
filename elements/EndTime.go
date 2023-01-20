package elements

// The EndTime element represents the end of a time span.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtime
import "time"

import "encoding/xml"

type EndTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndTime) SetForMarshal() {
	E.XMLName.Local = "t:EndTime"
}

func (E *EndTime) GetSchema() *Schema {
	return &SchemaTypes
}
