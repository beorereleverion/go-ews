package elements

// The SeekToConditionPageItemView element identifies the condition that is used to identify the end of a search, the starting index of a search, the maximum entries to return, and the search directions for a FindItem or FindConversation search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/seektoconditionpageitemview
type SeekToConditionPageItemView struct {
	// The Condition element specifies the condition that is used to identify the end of a search for a FindItem or a FindConversation operation.
	Condition *ConditionRestrictionType `xml:"t:Condition"`
	// The text value of the BasePoint attribute is the base point from where the search will start. A text value of Beginning indicates that the search will start at the beginning of the result set. A text value of End indicates that the search will start at the end of the result set.
	BasePoint *string `xml:"BasePoint,attr"`
	// The text value of the MaxEntriesReturned attribute is the maximum number of items that can be returned in a result set.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
}
