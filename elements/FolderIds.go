package elements

// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderids
import "encoding/xml"

type FolderIds struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (F *FolderIds) SetForMarshal() {
	F.XMLName.Local = "m:FolderIds"
}

func (F *FolderIds) GetSchema() *Schema {
	return &SchemaMessages
}
