package elements

// The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimesent
import "time"

import "encoding/xml"

type DateTimeSent struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTimeSent) SetForMarshal() {
	D.XMLName.Local = "t:DateTimeSent"
}

func (D *DateTimeSent) GetSchema() *Schema {
	return &SchemaTypes
}
