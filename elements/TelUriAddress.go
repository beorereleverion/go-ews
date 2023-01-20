package elements

// The TelUriAddress element contains the tel Uniform Resource Identifier (URI) for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/teluriaddress
import "encoding/xml"

type TelUriAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TelUriAddress) SetForMarshal() {
	T.XMLName.Local = "m:TelUriAddress"
}

func (T *TelUriAddress) GetSchema() *Schema {
	return &SchemaMessages
}
