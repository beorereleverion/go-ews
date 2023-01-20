package elements

// The UserIds element contains an array of delegate users to get or remove from a principal's mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userids
import "encoding/xml"

type UserIds struct {
	XMLName xml.Name

	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"UserId"`
}

func (U *UserIds) SetForMarshal() {
	U.XMLName.Local = "m:UserIds"
}

func (U *UserIds) GetSchema() *Schema {
	return &SchemaMessages
}
