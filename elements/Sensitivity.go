package elements

// The Sensitivity element indicates the sensitivity level of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sensitivity
import "encoding/xml"

type Sensitivity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Sensitivity) SetForMarshal() {
	S.XMLName.Local = "t:Sensitivity"
}

func (S *Sensitivity) GetSchema() *Schema {
	return &SchemaTypes
}
