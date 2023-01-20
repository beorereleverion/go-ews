package elements

// The Attendees element specifies the recipients of an invitation to a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendees
import "encoding/xml"

type Attendees struct {
	XMLName xml.Name

	// The EmailUser element specifies an email recipient.
	EmailUser *EmailUser `xml:"EmailUser"`
}

func (A *Attendees) SetForMarshal() {
	A.XMLName.Local = "t:Attendees"
}

func (A *Attendees) GetSchema() *Schema {
	return &SchemaTypes
}
