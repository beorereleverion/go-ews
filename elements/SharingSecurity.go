package elements

// The SharingSecurity element is used in the Simple Object Access Protocol (SOAP) header for calendar or contact sharing.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharingsecurity
import "encoding/xml"

type SharingSecurity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SharingSecurity) SetForMarshal() {
	S.XMLName.Local = "t:SharingSecurity"
}

func (S *SharingSecurity) GetSchema() *Schema {
	return &SchemaTypes
}
