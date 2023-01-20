package elements

// The UmEnabled element indicates whether Unified Messaging is enabled for an account.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/umenabled
import "encoding/xml"

type UmEnabled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	UmEnabledfalse bool = false
	// true
	UmEnabledtrue bool = true
)

func (U *UmEnabled) SetForMarshal() {
	U.XMLName.Local = "t:UmEnabled"
}

func (U *UmEnabled) GetSchema() *Schema {
	return &SchemaTypes
}
