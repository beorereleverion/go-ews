package elements

// The FolderChanges element represents a collection of changes for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderchanges
import "encoding/xml"

type FolderChanges struct {
	XMLName xml.Name

	// The FolderChange element represents a collection of changes to be performed on a single folder.
	FolderChange *FolderChange `xml:"FolderChange"`
}

func (F *FolderChanges) SetForMarshal() {
	F.XMLName.Local = "m:FolderChanges"
}

func (F *FolderChanges) GetSchema() *Schema {
	return &SchemaMessages
}
