package elements

// The IsFolderOwner element indicates whether a user is the owner of a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isfolderowner
import "encoding/xml"

type IsFolderOwner struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsFolderOwner) SetForMarshal() {
	I.XMLName.Local = "t:IsFolderOwner"
}

func (I *IsFolderOwner) GetSchema() *Schema {
	return &SchemaTypes
}
