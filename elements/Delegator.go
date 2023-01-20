package elements

// The Delegator element contains the name of the delegator who assigned the task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegator
import "encoding/xml"

type Delegator struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *Delegator) SetForMarshal() {
	D.XMLName.Local = "t:Delegator"
}

func (D *Delegator) GetSchema() *Schema {
	return &SchemaTypes
}
