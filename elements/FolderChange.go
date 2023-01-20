package elements

// The FolderChange element represents a collection of changes to be performed on a single folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderchange
import "encoding/xml"

type FolderChange struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
	// The Updates element contains a set of elements that define append, set, and delete changes to folder properties.
	Updates *UpdatesFolder `xml:"Updates"`
}

func (F *FolderChange) SetForMarshal() {
	F.XMLName.Local = "t:FolderChange"
}

func (F *FolderChange) GetSchema() *Schema {
	return &SchemaTypes
}
