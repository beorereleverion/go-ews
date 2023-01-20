package elements

// The Modify element indicates whether a client can modify a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modify
import "encoding/xml"

type Modify struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	Modifyfalse bool = false
	// true
	Modifytrue bool = true
)

func (M *Modify) SetForMarshal() {
	M.XMLName.Local = "t:Modify"
}

func (M *Modify) GetSchema() *Schema {
	return &SchemaTypes
}
