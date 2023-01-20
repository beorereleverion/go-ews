package elements

// The TooBigGroupAttendeeConflictData element represents an attendee that was resolved as a distribution list but the distribution list was too large to expand.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/toobiggroupattendeeconflictdata
import "encoding/xml"

type TooBigGroupAttendeeConflictData struct {
	XMLName xml.Name
	TEXT    struct{} `xml:",chardata"`
}

func (T *TooBigGroupAttendeeConflictData) SetForMarshal() {
	T.XMLName.Local = "t:TooBigGroupAttendeeConflictData"
}

func (T *TooBigGroupAttendeeConflictData) GetSchema() *Schema {
	return &SchemaTypes
}
