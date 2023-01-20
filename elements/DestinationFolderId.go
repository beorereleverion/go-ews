package elements

// The DestinationFolderId element indicates the destination folder for copy and move actions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/destinationfolderid
import "encoding/xml"

type DestinationFolderId struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (D *DestinationFolderId) SetForMarshal() {
	D.XMLName.Local = "t:DestinationFolderId"
}

func (D *DestinationFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
