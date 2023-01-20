package elements

// The OldItemId element contains the unique identifier of the item that was copied or moved.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/olditemid
import "encoding/xml"

type OldItemId struct {
	XMLName xml.Name

	// Contains a string that identifies a version of an item that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of an item is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies an item in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (O *OldItemId) SetForMarshal() {
	O.XMLName.Local = "t:OldItemId"
}

func (O *OldItemId) GetSchema() *Schema {
	return &SchemaTypes
}
