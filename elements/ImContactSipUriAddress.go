package elements

// The ImContactSipUriAddress element contains the SIP URI address of a contact that is added to an instant messaging (IM) group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imcontactsipuriaddress
import "encoding/xml"

type ImContactSipUriAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ImContactSipUriAddress) SetForMarshal() {
	I.XMLName.Local = "m:ImContactSipUriAddress"
}

func (I *ImContactSipUriAddress) GetSchema() *Schema {
	return &SchemaMessages
}
