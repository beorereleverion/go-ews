package elements

// The ArchiveSourceFolderId element specifies the Id of the source folder for the archive item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archivesourcefolderid
import "encoding/xml"

type ArchiveSourceFolderId struct {
	XMLName xml.Name

	// The AddressListId element specifies the identifier of an address list.
	AddressListId *AddressListId `xml:"AddressListId"`
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (A *ArchiveSourceFolderId) SetForMarshal() {
	A.XMLName.Local = "m:ArchiveSourceFolderId"
}

func (A *ArchiveSourceFolderId) GetSchema() *Schema {
	return &SchemaMessages
}
