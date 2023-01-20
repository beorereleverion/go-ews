package elements

// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemid
import "encoding/xml"

type ItemId struct {
	XMLName xml.Name

	// Identifies a specific version of an item. A ChangeKey is required for the following scenarios:  - The UpdateItem element requires a ChangeKey if the ConflictResolution attribute is set to AutoResolve. AutoResolve is a default value. If the ChangeKey attribute is not included, the response will return a ResponseCode value equal to ErrorChangeKeyRequired.  - The SendItem element requires a ChangeKey to test whether the attempted operation will act upon the most recent version of an item. If the ChangeKey attribute is not included in the ItemId or if the ChangeKey is empty, the response will return a ResponseCode value equal to ErrorStaleObject.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a specific item in the Exchange store. Id is case-sensitive; therefore, comparisons between Ids must be case-sensitive or binary.
	Id *string `xml:"Id,attr"`
}

func (I *ItemId) SetForMarshal() {
	I.XMLName.Local = "t:ItemId"
}

func (I *ItemId) GetSchema() *Schema {
	return &SchemaTypes
}
