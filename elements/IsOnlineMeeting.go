package elements

// The IsOnlineMeeting element indicates whether the meeting is online.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isonlinemeeting
import "encoding/xml"

type IsOnlineMeeting struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsOnlineMeeting) SetForMarshal() {
	I.XMLName.Local = "t:IsOnlineMeeting"
}

func (I *IsOnlineMeeting) GetSchema() *Schema {
	return &SchemaTypes
}
