package elements

// The CreateContents element indicates whether a client can create a contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createcontents
import "encoding/xml"

type CreateContents struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	CreateContentsfalse bool = false
	// true
	CreateContentstrue bool = true
)

func (C *CreateContents) SetForMarshal() {
	C.XMLName.Local = "t:CreateContents"
}

func (C *CreateContents) GetSchema() *Schema {
	return &SchemaTypes
}
