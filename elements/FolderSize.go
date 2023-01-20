package elements

// The FolderSize element describes the total size of all the contents of a managed folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldersize
import "encoding/xml"

type FolderSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (F *FolderSize) SetForMarshal() {
	F.XMLName.Local = "t:FolderSize"
}

func (F *FolderSize) GetSchema() *Schema {
	return &SchemaTypes
}
