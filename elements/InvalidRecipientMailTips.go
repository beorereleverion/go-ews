package elements

// The InvalidRecipient element indicates whether the recipient is invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/invalidrecipient-mailtips
import "encoding/xml"

type InvalidRecipientMailTips struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	InvalidRecipientMailTipsfalse bool = false
	// true
	InvalidRecipientMailTipstrue bool = true
)

func (I *InvalidRecipientMailTips) SetForMarshal() {
	I.XMLName.Local = "t:InvalidRecipient"
}

func (I *InvalidRecipientMailTips) GetSchema() *Schema {
	return &SchemaTypes
}
