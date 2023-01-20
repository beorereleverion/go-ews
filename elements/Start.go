package elements

// The Start element represents the start of a duration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/start
import "time"

import "encoding/xml"

type Start struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *Start) SetForMarshal() {
	S.XMLName.Local = "t:Start"
}

func (S *Start) GetSchema() *Schema {
	return &SchemaTypes
}
