package elements

// The End element represents the end of a duration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/end-ex15websvcsotherref
import "time"

import "encoding/xml"

type End struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *End) SetForMarshal() {
	E.XMLName.Local = "t:End"
}

func (E *End) GetSchema() *Schema {
	return &SchemaTypes
}
