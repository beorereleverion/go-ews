package elements

// The ExternalEmailAddress element contains the external email address of the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externalemailaddress
import "encoding/xml"

type ExternalEmailAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ExternalEmailAddress) SetForMarshal() {
	E.XMLName.Local = "t:ExternalEmailAddress"
}

func (E *ExternalEmailAddress) GetSchema() *Schema {
	return &SchemaTypes
}
