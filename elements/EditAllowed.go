package elements

// The EditAllowed element specifies whether Information Rights Management can be edited.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/editallowed
import "encoding/xml"

type EditAllowed struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	EditAllowedfalse bool = false
	// true
	EditAllowedtrue bool = true
)

func (E *EditAllowed) SetForMarshal() {
	E.XMLName.Local = "t:EditAllowed"
}

func (E *EditAllowed) GetSchema() *Schema {
	return &SchemaTypes
}
