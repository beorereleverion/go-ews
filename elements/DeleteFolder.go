package elements

// The DeleteFolder element defines a request to delete folders from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolder
import "encoding/xml"

type DeleteFolder struct {
	XMLName xml.Name

	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"FolderIds"`
	// Describes how a folder is deleted. This attribute is required.
	DeleteType *string `xml:"DeleteType,attr"`
}

func (D *DeleteFolder) SetForMarshal() {
	D.XMLName.Local = "m:DeleteFolder"
}

func (D *DeleteFolder) GetSchema() *Schema {
	return &SchemaMessages
}
