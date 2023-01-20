package elements

// The AttendeeLocation element specifies the location of an attendee for a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeelocation
import "encoding/xml"

type AttendeeLocation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *AttendeeLocation) SetForMarshal() {
	A.XMLName.Local = "t:AttendeeLocation"
}

func (A *AttendeeLocation) GetSchema() *Schema {
	return &SchemaTypes
}
