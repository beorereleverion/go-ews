package elements

// The GlobalImportance element contains the aggregated importance for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalimportance
import "encoding/xml"

type GlobalImportance struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *GlobalImportance) SetForMarshal() {
	G.XMLName.Local = "t:GlobalImportance"
}

func (G *GlobalImportance) GetSchema() *Schema {
	return &SchemaTypes
}
