package elements

// The GlobalParentFolderIds element specifies the identifiers of the global parent folders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalparentfolderids
import "encoding/xml"

type GlobalParentFolderIds struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (G *GlobalParentFolderIds) SetForMarshal() {
	G.XMLName.Local = "t:GlobalParentFolderIds"
}

func (G *GlobalParentFolderIds) GetSchema() *Schema {
	return &SchemaTypes
}
