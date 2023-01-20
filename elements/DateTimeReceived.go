package elements

// The DateTimeReceived element represents the date and time that an item in a mailbox was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimereceived
import "time"

import "encoding/xml"

type DateTimeReceived struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTimeReceived) SetForMarshal() {
	D.XMLName.Local = "t:DateTimeReceived"
}

func (D *DateTimeReceived) GetSchema() *Schema {
	return &SchemaTypes
}
