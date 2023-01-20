package elements

// The HiddenRecipient element indicates that the recipient was added by an organization policy that should be hidden from unprivileged users.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hiddenrecipient
import "encoding/xml"

type HiddenRecipient struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HiddenRecipientfalse bool = false
	// true
	HiddenRecipienttrue bool = true
)

func (H *HiddenRecipient) SetForMarshal() {
	H.XMLName.Local = "t:HiddenRecipient"
}

func (H *HiddenRecipient) GetSchema() *Schema {
	return &SchemaTypes
}
