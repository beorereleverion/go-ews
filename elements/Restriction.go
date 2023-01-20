package elements

// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/restriction
import "encoding/xml"

type Restriction struct {
	XMLName xml.Name

	// The And element represents a search expression that allows you to perform a Boolean AND operation between two or more search expressions. The result of the AND operation is true if all the search expressions contained within the And element are true.
	And *And `xml:"And"`
	// The Contains element represents a search expression that determines whether a given property contains the supplied constant string value.
	Contains *Contains `xml:"Contains"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"Excludes"`
	// The Exists element represents a search expression that returns true if the supplied property exists on an item.
	Exists *Exists `xml:"Exists"`
	// The IsEqualTo element represents a search expression that compares a property with either a constant value or another property and evaluates to true if they are equal.
	IsEqualTo *IsEqualTo `xml:"IsEqualTo"`
	// The IsGreaterThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater.
	IsGreaterThan *IsGreaterThan `xml:"IsGreaterThan"`
	// The IsGreaterThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is greater than or equal to the second.
	IsGreaterThanOrEqualTo *IsGreaterThanOrEqualTo `xml:"IsGreaterThanOrEqualTo"`
	// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
	IsLessThan *IsLessThan `xml:"IsLessThan"`
	// The IsLessThanOrEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than or equal to the second.
	IsLessThanOrEqualTo *IsLessThanOrEqualTo `xml:"IsLessThanOrEqualTo"`
	// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
	IsNotEqualTo *IsNotEqualTo `xml:"IsNotEqualTo"`
	// The Not element represents a search expression that negates the Boolean value of the search expression that it contains.
	Not *Not `xml:"Not"`
	// The Or element represents a search expression that performs a logical OR on the search expression that it contains. Or will return true if any of its children return true. Or must have two or more children.
	Or *Or `xml:"Or"`
	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"SearchExpression"`
}

func (R *Restriction) SetForMarshal() {
	R.XMLName.Local = "m:Restriction"
}

func (R *Restriction) GetSchema() *Schema {
	return &SchemaMessages
}
