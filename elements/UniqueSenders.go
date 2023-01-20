package elements

// The UniqueSenders element contains a list of all the senders of conversation items in the current folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquesenders
import "encoding/xml"

type UniqueSenders struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (U *UniqueSenders) SetForMarshal() {
	U.XMLName.Local = "t:UniqueSenders"
}

func (U *UniqueSenders) GetSchema() *Schema {
	return &SchemaTypes
}
