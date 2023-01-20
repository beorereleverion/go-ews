package elements

// The FolderIds element contains a list of folder identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderids-arrayoffolderidtype
import "encoding/xml"

type FolderIdsArrayOfFolderIdType struct {
	XMLName xml.Name

	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (F *FolderIdsArrayOfFolderIdType) SetForMarshal() {
	F.XMLName.Local = "t:FolderIds"
}

func (F *FolderIdsArrayOfFolderIdType) GetSchema() *Schema {
	return &SchemaTypes
}
