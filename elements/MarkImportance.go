package elements

// The MarkImportance element specifies the importance that is to be stamped on messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markimportance
import "encoding/xml"

type MarkImportance struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MarkImportance) SetForMarshal() {
	M.XMLName.Local = "m:MarkImportance"
}

func (M *MarkImportance) GetSchema() *Schema {
	return &SchemaMessages
}
