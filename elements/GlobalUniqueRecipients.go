package elements

// The GlobalUniqueRecipients element contains the recipient list of a conversation aggregated across a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaluniquerecipients
import "encoding/xml"

type GlobalUniqueRecipients struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (G *GlobalUniqueRecipients) SetForMarshal() {
	G.XMLName.Local = "t:GlobalUniqueRecipients"
}

func (G *GlobalUniqueRecipients) GetSchema() *Schema {
	return &SchemaTypes
}
