package elements

// The SubmitTime element represents the time at which the message was sent to the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/submittime
import "time"

import "encoding/xml"

type SubmitTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *SubmitTime) SetForMarshal() {
	S.XMLName.Local = "t:SubmitTime"
}

func (S *SubmitTime) GetSchema() *Schema {
	return &SchemaTypes
}
