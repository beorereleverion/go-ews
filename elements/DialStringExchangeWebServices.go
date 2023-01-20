package elements

// The DialString element represents the dial string of the telephone number that is called to play an item by telephone. This element is required.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dialstring-exchange-web-services
import "encoding/xml"

type DialStringExchangeWebServices struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DialStringExchangeWebServices) SetForMarshal() {
	D.XMLName.Local = "m:DialString"
}

func (D *DialStringExchangeWebServices) GetSchema() *Schema {
	return &SchemaMessages
}
