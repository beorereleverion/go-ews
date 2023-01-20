package elements

// The ManagedFolderId element contains the folder ID of the managed folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managedfolderid
import "encoding/xml"

type ManagedFolderId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *ManagedFolderId) SetForMarshal() {
	M.XMLName.Local = "t:ManagedFolderId"
}

func (M *ManagedFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
