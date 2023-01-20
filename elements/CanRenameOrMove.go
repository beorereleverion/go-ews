package elements

// The CanRenameOrMove element indicates whether a managed folder can be renamed or moved by the customer.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/canrenameormove
import "encoding/xml"

type CanRenameOrMove struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (C *CanRenameOrMove) SetForMarshal() {
	C.XMLName.Local = "t:CanRenameOrMove"
}

func (C *CanRenameOrMove) GetSchema() *Schema {
	return &SchemaTypes
}
