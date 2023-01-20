package elements

// The ContextFolderId element indicates the folder that is targeted for actions that use folders. This element must be present when copying, deleting, moving, and setting read state on conversation items in a target folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contextfolderid
import "encoding/xml"

type ContextFolderId struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (C *ContextFolderId) SetForMarshal() {
	C.XMLName.Local = "t:ContextFolderId"
}

func (C *ContextFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
