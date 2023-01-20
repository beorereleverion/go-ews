package elements

// The IsExternalMailbox element indicates whether the mailbox is external to the organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isexternalmailbox
import "encoding/xml"

type IsExternalMailbox struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsExternalMailboxfalse bool = false
	// true
	IsExternalMailboxtrue bool = true
)

func (I *IsExternalMailbox) SetForMarshal() {
	I.XMLName.Local = "t:IsExternalMailbox"
}

func (I *IsExternalMailbox) GetSchema() *Schema {
	return &SchemaTypes
}
