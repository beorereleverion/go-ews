package elements

// The State element represents the state of residence for a contact item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/state-ex15websvcsotherref
import "encoding/xml"

type State struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *State) SetForMarshal() {
	S.XMLName.Local = "t:State"
}

func (S *State) GetSchema() *Schema {
	return &SchemaTypes
}
