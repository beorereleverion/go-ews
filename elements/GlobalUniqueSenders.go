package elements

// The GlobalUniqueSender element contains a list of all the senders of conversation items in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaluniquesenders
import "encoding/xml"

type GlobalUniqueSenders struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (G *GlobalUniqueSenders) SetForMarshal() {
	G.XMLName.Local = "t:GlobalUniqueSenders"
}

func (G *GlobalUniqueSenders) GetSchema() *Schema {
	return &SchemaTypes
}
