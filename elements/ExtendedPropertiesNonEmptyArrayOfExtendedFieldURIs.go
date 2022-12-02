package elements

// The ExtendedProperties element contains the extended properties used for the Unified Contact Store operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-nonemptyarrayofextendedfielduris
type ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs struct {
	// The ExtendedProperty element specifies an extended property for the Unified Contact Store.
	ExtendedProperty *ExtendedPropertyPathToExtendedFieldType `xml:"t:ExtendedProperty"`
}
