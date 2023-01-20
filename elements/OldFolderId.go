package elements

// The OldFolderId element contains the original identifier of a folder that was moved or copied.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oldfolderid
import "encoding/xml"

type OldFolderId struct {
	XMLName xml.Name

	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (O *OldFolderId) SetForMarshal() {
	O.XMLName.Local = "t:OldFolderId"
}

func (O *OldFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
