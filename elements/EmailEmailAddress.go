package elements

// The Email element identifies an attendee to a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/email-emailaddress
import "encoding/xml"

type EmailEmailAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EmailEmailAddress) SetForMarshal() {
	E.XMLName.Local = "t:Email"
}

func (E *EmailEmailAddress) GetSchema() *Schema {
	return &SchemaTypes
}
