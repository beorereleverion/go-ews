package elements

// The FolderClass element represents the folder class for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderclass
import "encoding/xml"

type FolderClass struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FolderClass) SetForMarshal() {
	F.XMLName.Local = "t:FolderClass"
}

func (F *FolderClass) GetSchema() *Schema {
	return &SchemaTypes
}
