package elements

// The Timestamp element represents the timestamp of a mailbox event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timestamp
import "time"

import "encoding/xml"

type TimeStamp struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (T *TimeStamp) SetForMarshal() {
	T.XMLName.Local = "t:TimeStamp"
}

func (T *TimeStamp) GetSchema() *Schema {
	return &SchemaTypes
}
