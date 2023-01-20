package elements

// The IsReadReceipt element indicates whether incoming messages must be read receipts in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isreadreceipt
import "encoding/xml"

type IsReadReceipt struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsReadReceiptfalse bool = false
	// true
	IsReadReceipttrue bool = true
)

func (I *IsReadReceipt) SetForMarshal() {
	I.XMLName.Local = "m:IsReadReceipt"
}

func (I *IsReadReceipt) GetSchema() *Schema {
	return &SchemaMessages
}
