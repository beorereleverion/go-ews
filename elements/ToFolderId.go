package elements

// The ToFolderId element represents the destination folder for a copied or moved item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tofolderid
import "encoding/xml"

type ToFolderId struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (T *ToFolderId) SetForMarshal() {
	T.XMLName.Local = "m:ToFolderId"
}

func (T *ToFolderId) GetSchema() *Schema {
	return &SchemaMessages
}
