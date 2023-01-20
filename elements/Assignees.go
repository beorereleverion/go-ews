package elements

// The Assignees element specifies the people to whom a task is assigned.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assignees
import "encoding/xml"

type Assignees struct {
	XMLName xml.Name

	// The Name element represents the display name of the mailbox user.
	Name *NameEmailAddress `xml:"Name"`
	// The UserId element specifies the user identifier of an email user.
	UserId *UserIdstring `xml:"UserId"`
}

func (A *Assignees) SetForMarshal() {
	A.XMLName.Local = "t:Assignees"
}

func (A *Assignees) GetSchema() *Schema {
	return &SchemaTypes
}
