package elements

// The UpdateFolder element represents the operation that is used to update properties for a specified folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolder
import "encoding/xml"

type UpdateFolder struct {
	XMLName xml.Name

	// The FolderChanges element represents a collection of changes for a folder.
	FolderChanges *FolderChanges `xml:"FolderChanges"`
}

func (U *UpdateFolder) SetForMarshal() {
	U.XMLName.Local = "m:UpdateFolder"
}

func (U *UpdateFolder) GetSchema() *Schema {
	return &SchemaMessages
}
