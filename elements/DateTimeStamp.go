package elements

// The DateTimeStamp element indicates the date and time that an instance of a calendar object was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimestamp
import "time"

import "encoding/xml"

type DateTimeStamp struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTimeStamp) SetForMarshal() {
	D.XMLName.Local = "t:DateTimeStamp"
}

func (D *DateTimeStamp) GetSchema() *Schema {
	return &SchemaTypes
}
