package elements

// The Condition element specifies the condition that is used to identify the end of a search for a FindItem or a FindConversation operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/condition-restrictiontype
type ConditionRestrictionType struct {
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression"`
}
