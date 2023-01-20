package elements

// The CopyToFolder element specifies the identifier of the folder that email items can be copied to.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copytofolder
import "encoding/xml"

type CopyToFolder struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (C *CopyToFolder) SetForMarshal() {
	C.XMLName.Local = "m:CopyToFolder"
}

func (C *CopyToFolder) GetSchema() *Schema {
	return &SchemaMessages
}
