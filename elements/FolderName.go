package elements

// The FolderName element identifies a single managed custom folder to add to a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldername
import "encoding/xml"

type FolderName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FolderName) SetForMarshal() {
	F.XMLName.Local = "t:FolderName"
}

func (F *FolderName) GetSchema() *Schema {
	return &SchemaTypes
}
