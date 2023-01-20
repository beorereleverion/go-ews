package elements

// The SyncFolderId element represents the folder that contains the items to synchronize.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderid
import "encoding/xml"

type SyncFolderId struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (S *SyncFolderId) SetForMarshal() {
	S.XMLName.Local = "m:SyncFolderId"
}

func (S *SyncFolderId) GetSchema() *Schema {
	return &SchemaMessages
}
