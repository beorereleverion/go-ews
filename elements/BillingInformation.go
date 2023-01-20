package elements

// The BillingInformation element holds billing information for a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/billinginformation
import "encoding/xml"

type BillingInformation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BillingInformation) SetForMarshal() {
	B.XMLName.Local = "t:BillingInformation"
}

func (B *BillingInformation) GetSchema() *Schema {
	return &SchemaTypes
}
