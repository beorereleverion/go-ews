package elements

// The ForwardAllowed element specifies whether forwarding emails is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardallowed
import "encoding/xml"

type ForwardAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ForwardAllowedfalse bool = false
	// true
	ForwardAllowedtrue bool = true
)

func (F *ForwardAllowed) SetForMarshal() {
	F.XMLName.Local = "t:ForwardAllowed"
}

func (F *ForwardAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
