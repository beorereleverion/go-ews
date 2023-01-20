package elements

// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortorder
import "encoding/xml"

type SortOrder struct {
	XMLName xml.Name

	// The FieldOrder element represents a single field by which to sort results and indicates the direction for the sort.
	FieldOrder *FieldOrder `xml:"FieldOrder"`
}

func (S *SortOrder) SetForMarshal() {
	S.XMLName.Local = "m:SortOrder"
}

func (S *SortOrder) GetSchema() *Schema {
	return &SchemaMessages
}
