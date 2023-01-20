package elements

// The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeitem
import "encoding/xml"

type RemoveItem struct {
	XMLName xml.Name

	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"ReferenceItemId"`
	// Represents the name of the RemoveItem reply object class as an English string.
	ObjectName *string `xml:"ObjectName,attr"`
}

func (R *RemoveItem) SetForMarshal() {
	R.XMLName.Local = "t:RemoveItem"
}

func (R *RemoveItem) GetSchema() *Schema {
	return &SchemaTypes
}
