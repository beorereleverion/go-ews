package elements

// The ExportItems element represents a request to export items from a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportitems
import "encoding/xml"

type ExportItems struct {
	XMLName xml.Name

	// The ItemIds element contains an array of item identifiers that identify the items to export from a mailbox.
	ItemIds *ItemIdsNonEmptyArrayOfItemIdsType `xml:"ItemIds"`
}

func (E *ExportItems) SetForMarshal() {
	E.XMLName.Local = "m:ExportItems"
}

func (E *ExportItems) GetSchema() *Schema {
	return &SchemaMessages
}
