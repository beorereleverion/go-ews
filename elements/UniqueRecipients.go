package elements

// The UniqueRecipients element contains the recipient list of a conversation in a particular folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquerecipients
import "encoding/xml"

type UniqueRecipients struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (U *UniqueRecipients) SetForMarshal() {
	U.XMLName.Local = "m:UniqueRecipients"
}

func (U *UniqueRecipients) GetSchema() *Schema {
	return &SchemaMessages
}
