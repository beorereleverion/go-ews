package elements

// The CanCreateItems element indicates whether a user has permission to create items in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/cancreateitems
import "encoding/xml"

type CanCreateItems struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (C *CanCreateItems) SetForMarshal() {
	C.XMLName.Local = "t:CanCreateItems"
}

func (C *CanCreateItems) GetSchema() *Schema {
	return &SchemaTypes
}
