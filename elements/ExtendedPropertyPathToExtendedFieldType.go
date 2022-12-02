package elements

// The ExtendedProperty element specifies an extended property for the Unified Contact Store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperty-pathtoextendedfieldtype
type ExtendedPropertyPathToExtendedFieldType struct {
	// Indicates the distinguished property set identifier. This attribute is optional.
	DistinguishedPropertySetId *string `xml:"DistinguishedPropertySetId,attr"`
	// Indicates the field Uniform Resource Identifier (URI). This attribute is required. For possible values, see the FieldURI element.
	FieldURI *string `xml:"FieldURI,attr"`
	// Integer that indicates the property identifier. This attribute is optional.
	PropertyId *string `xml:"PropertyId,attr"`
	// String that indicates the property name. This attribute is optional.
	PropertyName *string `xml:"PropertyName,attr"`
	// Indicates the GUID property set identifier. This attribute is optional.
	PropertySetId *string `xml:"PropertySetId,attr"`
	// Represents the property tag minus the type part.There are two options for representation:  - Hexadecimal: 0x3fa4  - Decimal: 0-65535  This attribute is optional.
	PropertyTag *string `xml:"PropertyTag,attr"`
	// Indicates the property type. This attribute is required.
	PropertyType *string `xml:"PropertyType,attr"`
}
