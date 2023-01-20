package elements

// The UserEnabledExtensions element lists the enabled apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userenabledextensions
import "encoding/xml"

type UserEnabledExtensions struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (U *UserEnabledExtensions) SetForMarshal() {
	U.XMLName.Local = "t:UserEnabledExtensions"
}

func (U *UserEnabledExtensions) GetSchema() *Schema {
	return &SchemaTypes
}
