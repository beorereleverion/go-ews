package elements

// The OriginalStart element represents the original start time of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalstart
import "time"

import "encoding/xml"

type OriginalStart struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (O *OriginalStart) SetForMarshal() {
	O.XMLName.Local = "t:OriginalStart"
}

func (O *OriginalStart) GetSchema() *Schema {
	return &SchemaTypes
}
