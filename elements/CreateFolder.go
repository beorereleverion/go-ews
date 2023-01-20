package elements

// The CreateFolder element defines a request to create a folder in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolder
import "encoding/xml"

type CreateFolder struct {
	XMLName xml.Name

	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"Folders"`
	// The ParentFolderId element identifies the folder in which a new folder is created or the folder to search for the FindConversation operation.
	ParentFolderId *ParentFolderIdTargetFolderIdType `xml:"ParentFolderId"`
}

func (C *CreateFolder) SetForMarshal() {
	C.XMLName.Local = "m:CreateFolder"
}

func (C *CreateFolder) GetSchema() *Schema {
	return &SchemaMessages
}
