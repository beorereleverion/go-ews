package elements

// The StartTime element represents the start of a time span.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime
import "time"

import "encoding/xml"

type StartTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartTime) SetForMarshal() {
	S.XMLName.Local = "t:StartTime"
}

func (S *StartTime) GetSchema() *Schema {
	return &SchemaTypes
}
