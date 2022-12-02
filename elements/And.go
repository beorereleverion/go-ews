package elements

// The And element represents a search expression that allows you to perform a Boolean AND operation between two or more search expressions. The result of the AND operation is true if all the search expressions contained within the And element are true.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/and
type And struct {
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"t:Contains"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"t:Exists"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"t:IsEqualTo"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"t:IsGreaterThan"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"t:IsGreaterThanOrEqualTo"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"t:IsLessThan"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"t:IsLessThanOrEqualTo"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"t:IsNotEqualTo"`
	// The Not element represents a search expression that negates the Boolean value of the search expression that it contains.
	Not *Not `xml:"t:Not"`
	// The Or element represents a search expression that performs a logical OR on the search expression that it contains. Or will return true if any of its children return true. Or must have two or more children.
	Or *Or `xml:"t:Or"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"t:SearchExpression"`
}
