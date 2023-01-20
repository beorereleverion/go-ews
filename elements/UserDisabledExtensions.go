package elements

// The UserDisabledExtensions element lists the disabled apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userdisabledextensions
import "encoding/xml"

type UserDisabledExtensions struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (U *UserDisabledExtensions) SetForMarshal() {
	U.XMLName.Local = "t:UserDisabledExtensions"
}

func (U *UserDisabledExtensions) GetSchema() *Schema {
	return &SchemaTypes
}
