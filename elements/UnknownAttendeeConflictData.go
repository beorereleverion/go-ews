package elements

// The UnknownAttendeeConflictData element represents an unresolvable attendee or an attendee that is not a user, distribution list, or contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unknownattendeeconflictdata
import "encoding/xml"

type UnknownAttendeeConflictData struct {
	XMLName xml.Name
	TEXT    struct{} `xml:",chardata"`
}

func (U *UnknownAttendeeConflictData) SetForMarshal() {
	U.XMLName.Local = "t:UnknownAttendeeConflictData"
}

func (U *UnknownAttendeeConflictData) GetSchema() *Schema {
	return &SchemaTypes
}
