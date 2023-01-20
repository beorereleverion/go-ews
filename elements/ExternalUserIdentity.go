package elements

// The ExternalUserIdentity element identifies an external delegate user or an external user who has folder access permissions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externaluseridentity
import "encoding/xml"

type ExternalUserIdentity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ExternalUserIdentity) SetForMarshal() {
	E.XMLName.Local = "t:ExternalUserIdentity"
}

func (E *ExternalUserIdentity) GetSchema() *Schema {
	return &SchemaTypes
}
