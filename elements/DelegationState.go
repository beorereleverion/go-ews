package elements

// The DelegationState element represents the status of a delegated task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegationstate
import "encoding/xml"

type DelegationState struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DelegationState) SetForMarshal() {
	D.XMLName.Local = "t:DelegationState"
}

func (D *DelegationState) GetSchema() *Schema {
	return &SchemaTypes
}
