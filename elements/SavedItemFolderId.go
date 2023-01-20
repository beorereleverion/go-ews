package elements

// The SavedItemFolderId element identifies the target folder for operations that update, send, and create items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/saveditemfolderid
import "encoding/xml"

type SavedItemFolderId struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (S *SavedItemFolderId) SetForMarshal() {
	S.XMLName.Local = "m:SavedItemFolderId"
}

func (S *SavedItemFolderId) GetSchema() *Schema {
	return &SchemaMessages
}
