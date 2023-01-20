package elements

// The IsFolderContact element indicates whether a user is a contact for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isfoldercontact
import "encoding/xml"

type IsFolderContact struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsFolderContact) SetForMarshal() {
	I.XMLName.Local = "t:IsFolderContact"
}

func (I *IsFolderContact) GetSchema() *Schema {
	return &SchemaTypes
}
