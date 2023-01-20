package elements

// The Address element represents the e-mail address of the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/address-string
import "encoding/xml"

type Addressstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *Addressstring) SetForMarshal() {
	A.XMLName.Local = "t:Address"
}

func (A *Addressstring) GetSchema() *Schema {
	return &SchemaTypes
}
