package elements

// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isdeliveryreceiptrequested
import "encoding/xml"

type IsDeliveryReceiptRequested struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsDeliveryReceiptRequested) SetForMarshal() {
	I.XMLName.Local = "t:IsDeliveryReceiptRequested"
}

func (I *IsDeliveryReceiptRequested) GetSchema() *Schema {
	return &SchemaTypes
}
