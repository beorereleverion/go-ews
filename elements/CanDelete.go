package elements

// The CanDelete element indicates whether a managed folder can be deleted by a customer.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/candelete
import "encoding/xml"

type CanDelete struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (C *CanDelete) SetForMarshal() {
	C.XMLName.Local = "t:CanDelete"
}

func (C *CanDelete) GetSchema() *Schema {
	return &SchemaTypes
}
