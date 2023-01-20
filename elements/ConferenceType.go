package elements

// The ConferenceType element describes the type of conferencing that is performed with a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conferencetype
import "encoding/xml"

type ConferenceType struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *ConferenceType) SetForMarshal() {
	C.XMLName.Local = "t:ConferenceType"
}

func (C *ConferenceType) GetSchema() *Schema {
	return &SchemaTypes
}
