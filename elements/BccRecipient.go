package elements

// The BccRecipient element represents a recipient to receive a blind carbon copy (Bcc) of an e-mail message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bccrecipient
import "encoding/xml"

type BccRecipient struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	BccRecipientfalse bool = false
	// true
	BccRecipienttrue bool = true
)

func (B *BccRecipient) SetForMarshal() {
	B.XMLName.Local = "t:BccRecipient"
}

func (B *BccRecipient) GetSchema() *Schema {
	return &SchemaTypes
}
