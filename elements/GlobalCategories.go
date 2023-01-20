package elements

// The GlobalCategories element contains the category list for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalcategories
import "encoding/xml"

type GlobalCategories struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (G *GlobalCategories) SetForMarshal() {
	G.XMLName.Local = "t:GlobalCategories"
}

func (G *GlobalCategories) GetSchema() *Schema {
	return &SchemaTypes
}
