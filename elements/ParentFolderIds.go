package elements

// The ParentFolderIds element identifies folders for the FindItem and FindFolder operations to search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderids
import "encoding/xml"

type ParentFolderIds struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (P *ParentFolderIds) SetForMarshal() {
	P.XMLName.Local = "m:ParentFolderIds"
}

func (P *ParentFolderIds) GetSchema() *Schema {
	return &SchemaMessages
}
