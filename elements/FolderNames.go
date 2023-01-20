package elements

// The FolderNames element contains an array of named managed folders to add to a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldernames
import "encoding/xml"

type FolderNames struct {
	XMLName xml.Name

	// The FolderName element identifies a single managed custom folder to add to a mailbox.
	FolderName *FolderName `xml:"FolderName"`
}

func (F *FolderNames) SetForMarshal() {
	F.XMLName.Local = "m:FolderNames"
}

func (F *FolderNames) GetSchema() *Schema {
	return &SchemaMessages
}
