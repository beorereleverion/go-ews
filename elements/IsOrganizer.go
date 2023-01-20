package elements

// The IsOrganizer element specifies a Boolean value that indicates whether this person is the organizer of the meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isorganizer
import "encoding/xml"

type IsOrganizer struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsOrganizerfalse bool = false
	// true
	IsOrganizertrue bool = true
)

func (I *IsOrganizer) SetForMarshal() {
	I.XMLName.Local = "t:IsOrganizer"
}

func (I *IsOrganizer) GetSchema() *Schema {
	return &SchemaTypes
}
