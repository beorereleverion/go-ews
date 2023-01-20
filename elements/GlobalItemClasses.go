package elements

// The GlobalItemClasses element contains a list of item classes that represents all the item classes of the conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalitemclasses
import "encoding/xml"

type GlobalItemClasses struct {
	XMLName xml.Name

	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
}

func (G *GlobalItemClasses) SetForMarshal() {
	G.XMLName.Local = "t:GlobalItemClasses"
}

func (G *GlobalItemClasses) GetSchema() *Schema {
	return &SchemaTypes
}
