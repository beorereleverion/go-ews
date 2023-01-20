package elements

// The End element specifies the changes made to a meeting end time when a meeting update occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/end-changehighlightstype
import "time"

import "encoding/xml"

type EndChangeHighlightsType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndChangeHighlightsType) SetForMarshal() {
	E.XMLName.Local = "t:End"
}

func (E *EndChangeHighlightsType) GetSchema() *Schema {
	return &SchemaTypes
}
