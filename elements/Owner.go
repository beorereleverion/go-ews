package elements

// The Owner element represents the owner of a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/owner
import "encoding/xml"

type Owner struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (O *Owner) SetForMarshal() {
	O.XMLName.Local = "t:Owner"
}

func (O *Owner) GetSchema() *Schema {
	return &SchemaTypes
}
