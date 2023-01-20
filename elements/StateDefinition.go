package elements

// The StateDefinition element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/statedefinition
import "encoding/xml"

type StateDefinition struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (S *StateDefinition) SetForMarshal() {
	S.XMLName.Local = "t:StateDefinition"
}

func (S *StateDefinition) GetSchema() *Schema {
	return &SchemaTypes
}
