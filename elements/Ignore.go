package elements

// The Ignore element identifies items to skip during synchronization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ignore
import "encoding/xml"

type Ignore struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (I *Ignore) SetForMarshal() {
	I.XMLName.Local = "m:Ignore"
}

func (I *Ignore) GetSchema() *Schema {
	return &SchemaMessages
}
