package elements

// The ItemChanges element contains an array of ItemChange elements that identify items and the updates to apply to the items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchanges
import "encoding/xml"

type ItemChanges struct {
	XMLName xml.Name

	// The ItemChange element contains an item identifier and the updates to apply to the item.
	ItemChange *ItemChange `xml:"ItemChange"`
}

func (I *ItemChanges) SetForMarshal() {
	I.XMLName.Local = "m:ItemChanges"
}

func (I *ItemChanges) GetSchema() *Schema {
	return &SchemaMessages
}
