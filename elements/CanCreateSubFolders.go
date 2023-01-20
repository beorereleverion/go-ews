package elements

// The CanCreateSubFolders element indicates whether a user has permission to create subfolders in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/cancreatesubfolders
import "encoding/xml"

type CanCreateSubFolders struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (C *CanCreateSubFolders) SetForMarshal() {
	C.XMLName.Local = "t:CanCreateSubFolders"
}

func (C *CanCreateSubFolders) GetSchema() *Schema {
	return &SchemaTypes
}
