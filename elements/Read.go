package elements

// The Read element indicates whether a client can read a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/read
import "encoding/xml"

type Read struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	Readfalse bool = false
	// true
	Readtrue bool = true
)

func (R *Read) SetForMarshal() {
	R.XMLName.Local = "t:Read"
}

func (R *Read) GetSchema() *Schema {
	return &SchemaTypes
}
