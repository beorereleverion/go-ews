package elements

// The FolderId element contains the identifier and change key of a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderid
import "encoding/xml"

type FolderId struct {
	XMLName xml.Name

	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (F *FolderId) SetForMarshal() {
	F.XMLName.Local = "t:FolderId"
}

func (F *FolderId) GetSchema() *Schema {
	return &SchemaTypes
}
