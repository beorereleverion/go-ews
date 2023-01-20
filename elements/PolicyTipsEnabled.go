package elements

// The PolicyTipsEnabled element indicates whether policy tips are enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/policytipsenabled
import "encoding/xml"

type PolicyTipsEnabled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	PolicyTipsEnabledfalse bool = false
	// true
	PolicyTipsEnabledtrue bool = true
)

func (P *PolicyTipsEnabled) SetForMarshal() {
	P.XMLName.Local = "t:PolicyTipsEnabled"
}

func (P *PolicyTipsEnabled) GetSchema() *Schema {
	return &SchemaTypes
}
