package elements

// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isreadreceiptrequested
import "encoding/xml"

type IsReadReceiptRequested struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsReadReceiptRequested) SetForMarshal() {
	I.XMLName.Local = "t:IsReadReceiptRequested"
}

func (I *IsReadReceiptRequested) GetSchema() *Schema {
	return &SchemaTypes
}
