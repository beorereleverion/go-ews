package elements

// The ExtendedProperties element contains the extended properties used for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-arrayofextendedpropertyattributedvaluetype
type ExtendedPropertiesArrayOfExtendedPropertyAttributedValueType struct {
	// The ExtendedPropertyAttributedValue element specifies extended properties for a persona.
	ExtendedPropertyAttributedValue *ExtendedPropertyAttributedValue `xml:"t:ExtendedPropertyAttributedValue"`
}
