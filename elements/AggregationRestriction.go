package elements

// The AggregationRestriction element specifies a value that is applied to a set of Persona properties resulting from a FindPeople request and filters the result according to the specified restriction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/aggregationrestriction
type AggregationRestriction struct {
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression"`
}
