package elements

// The CreateFolderPath element is used to create a folder path and includes a parent folder Id and a relative folder path.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderpath
import "encoding/xml"

type CreateFolderPath struct {
	XMLName xml.Name

	// The ParentFolderId element identifies the folder in which a new folder is created or the folder to search for the FindConversation operation.
	ParentFolderId *ParentFolderIdTargetFolderIdType `xml:"ParentFolderId"`
	// The RelativeFolderPath element contains an array of folders that indicate the relative folder path of the folder path to be created.
	RelativeFolderPath *RelativeFolderPath `xml:"RelativeFolderPath"`
}

func (C *CreateFolderPath) SetForMarshal() {
	C.XMLName.Local = "m:CreateFolderPath"
}

func (C *CreateFolderPath) GetSchema() *Schema {
	return &SchemaMessages
}
