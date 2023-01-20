package elements

// The OldParentFolderId element contains the identifier of the parent folder of an item or folder that was copied or moved.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oldparentfolderid
import "encoding/xml"

type OldParentFolderId struct {
	XMLName xml.Name

	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (O *OldParentFolderId) SetForMarshal() {
	O.XMLName.Local = "t:OldParentFolderId"
}

func (O *OldParentFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
