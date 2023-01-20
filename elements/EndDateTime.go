package elements

// The EndDateTime element specifies the end date and time for a rule or a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddatetime
import "time"

import "encoding/xml"

type EndDateTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndDateTime) SetForMarshal() {
	E.XMLName.Local = "m:EndDateTime"
}

func (E *EndDateTime) GetSchema() *Schema {
	return &SchemaMessages
}
