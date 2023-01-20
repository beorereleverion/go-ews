package elements

// The CreateAssociated element indicates whether a client can create an associated contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createassociated
import "encoding/xml"

type CreateAssociated struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	CreateAssociatedfalse bool = false
	// true
	CreateAssociatedtrue bool = true
)

func (C *CreateAssociated) SetForMarshal() {
	C.XMLName.Local = "t:CreateAssociated"
}

func (C *CreateAssociated) GetSchema() *Schema {
	return &SchemaTypes
}
