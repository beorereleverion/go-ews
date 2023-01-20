package elements

// The ProposedStart element specifies the proposed start time of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart
import "time"

import "encoding/xml"

type ProposedStart struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedStart) SetForMarshal() {
	P.XMLName.Local = "t:ProposedStart"
}

func (P *ProposedStart) GetSchema() *Schema {
	return &SchemaTypes
}
