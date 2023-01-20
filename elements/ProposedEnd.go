package elements

// The ProposedEnd element specifies the proposed end time of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend
import "time"

import "encoding/xml"

type ProposedEnd struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedEnd) SetForMarshal() {
	P.XMLName.Local = "t:ProposedEnd"
}

func (P *ProposedEnd) GetSchema() *Schema {
	return &SchemaTypes
}
