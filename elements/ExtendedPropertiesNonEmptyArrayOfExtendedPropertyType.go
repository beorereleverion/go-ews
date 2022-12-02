package elements

// The ExtendedProperties element specifies an array of additional properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-nonemptyarrayofextendedpropertytype
type ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType struct {
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty"`
}
