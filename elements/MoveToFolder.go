package elements

// The MoveToFolder element specifies the identifier of the folder to which email items can be moved.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/movetofolder
import "encoding/xml"

type MoveToFolder struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (M *MoveToFolder) SetForMarshal() {
	M.XMLName.Local = "m:MoveToFolder"
}

func (M *MoveToFolder) GetSchema() *Schema {
	return &SchemaMessages
}
