package elements

// The SentTime element specifies the time at which the item was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttime
import "time"

import "encoding/xml"

type SentTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *SentTime) SetForMarshal() {
	S.XMLName.Local = "t:SentTime"
}

func (S *SentTime) GetSchema() *Schema {
	return &SchemaTypes
}
