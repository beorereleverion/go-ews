package elements

// The SortOrder element specifies the sort order used for the result of a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortorder-conversationnodesortorder
import "encoding/xml"

type SortOrderConversationNodeSortOrder struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// DateOrderAscending
	SortOrderConversationNodeSortOrderDateOrderAscending string = `DateOrderAscending`
	// DateOrderDescending
	SortOrderConversationNodeSortOrderDateOrderDescending string = `DateOrderDescending`
	// TreeOrderAscending
	SortOrderConversationNodeSortOrderTreeOrderAscending string = `TreeOrderAscending`
	// TreeOrderDescending
	SortOrderConversationNodeSortOrderTreeOrderDescending string = `TreeOrderDescending`
)

func (S *SortOrderConversationNodeSortOrder) SetForMarshal() {
	S.XMLName.Local = "m:SortOrder"
}

func (S *SortOrderConversationNodeSortOrder) GetSchema() *Schema {
	return &SchemaMessages
}
