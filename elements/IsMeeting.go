package elements

// The IsMeeting element indicates whether the calendar item is a meeting or an appointment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeeting
import "encoding/xml"

type IsMeeting struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsMeeting) SetForMarshal() {
	I.XMLName.Local = "t:IsMeeting"
}

func (I *IsMeeting) GetSchema() *Schema {
	return &SchemaTypes
}
