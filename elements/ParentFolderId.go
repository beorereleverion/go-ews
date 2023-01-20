package elements

// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid
import "encoding/xml"

type ParentFolderId struct {
	XMLName xml.Name

	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (P *ParentFolderId) SetForMarshal() {
	P.XMLName.Local = "t:ParentFolderId"
}

func (P *ParentFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
