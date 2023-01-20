package elements

// The IsFolderVisible element indicates whether a user can view a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isfoldervisible
import "encoding/xml"

type IsFolderVisible struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsFolderVisible) SetForMarshal() {
	I.XMLName.Local = "t:IsFolderVisible"
}

func (I *IsFolderVisible) GetSchema() *Schema {
	return &SchemaTypes
}
