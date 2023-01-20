package elements

// The Guid element specifies the globally unique identifier of the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/guid-ex15websvcsotherref
import "encoding/xml"

type Guid struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *Guid) SetForMarshal() {
	G.XMLName.Local = "t:Guid"
}

func (G *Guid) GetSchema() *Schema {
	return &SchemaTypes
}
