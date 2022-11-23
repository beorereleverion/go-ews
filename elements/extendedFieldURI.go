package elements

// The ExtendedFieldURI element identifies an extended MAPI property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedfielduri
type ExtendedFieldURI struct {
	// Defines the well-known property set IDs for extended MAPI properties.
	// If this attribute is used, the PropertySetId and PropertyTag attributes cannot be used. This attribute must be used with either the PropertyId or PropertyName attribute, and the PropertyType attribute.
	// The DistinguishedPropertySetId Attribute table later in this topic lists the possible values for this attribute.5
	// This attribute is optional.
	DistinguishedPropertySetId *DistinguishedPropertySetId `xml:"DistinguishedPropertySetId,attr"`
	// Identifies a MAPI extended property set or namespace by its identifying GUID.
	// If this attribute is used, the DistinguishedPropertySetId and PropertyTag attribute cannot be used. This attribute must be used with either the PropertyId or PropertyName attribute, and the PropertyType attribute.
	// This attribute is optional.
	PropertySetId *string `xml:"PropertySetId,attr"`
	// Identifies the property tag without the type part of the tag. The PropertyTag can be represented as either a hexadecimal or a short integer.
	// The range between 0x8000 and 0xFFFE represents the custom range of properties. When a mailbox database encounters a custom property for the first time, it assigns that custom property a property tag within the custom property range of 0x8000-0xFFFE. A given custom property tag will most likely differ across databases. Therefore, a custom property request by property tag can return different properties on different databases. The use of the PropertyTag attribute is prohibited for custom properties. Instead, use the PropertySetId attribute and the PropertyName or PropertyId attribute.
	// IMPORTANT: Access any custom property between 0x8000 and 0xFFFE by using the GUID + name/ID. If the PropertyTag attribute is used, the DistinguishedPropertySetId, PropertySetId, PropertyName, and PropertyId attributes cannot be used.
	// This attribute is optional.
	// NOTE: You cannot use a property tag attribute for properties within the custom range 0x8000-0xFFFE. You must use a named property in this case.
	PropertyTag *string `xml:"PropertyTag,attr"`
	// Identifies an extended property by its name. This property must be coupled with either DistinguishedPropertySetId or PropertySetId.
	// If this attribute is used, the PropertyId and PropertyTag attributes cannot be used.
	// This attribute is optional.
	PropertyName *string `xml:"PropertyName,attr"`
	// Identifies an extended property by its dispatch ID. The dispatch ID can be identified in either decimal or hexadecimal formats. This property must be coupled with either DistinguishedPropertySetId or PropertySetId.
	// If this attribute is used, the PropertyName and PropertyTag attributes cannot be used.
	// This attribute is optional.
	PropertyId *string `xml:"PropertyId,attr"`
	// Represents the property type of a property tag. This corresponds to the least significant word in a property tag.
	// The PropertyType Attribute table later in this topic contains the possible values for this attribute.
	// This attribute is required.
	PropertyType *PropertyType `xml:"PropertyType,attr"`
}

type DistinguishedPropertySetId string

const (
	// Identifies the address property set ID by name.
	DistinguishedPropertySetIdAddress DistinguishedPropertySetId = `Address`
	// Identifies the appointment property set ID by name.
	DistinguishedPropertySetIdAppointment DistinguishedPropertySetId = `Appointment`
	// Identifies the calendar assistant property set ID by name.
	DistinguishedPropertySetIdCalendarAssistant DistinguishedPropertySetId = `CalendarAssistant`
	// Identifies the common property set ID by name.
	DistinguishedPropertySetIdCommon DistinguishedPropertySetId = `Common`
	// Identifies the Internet headers property set ID by name.
	DistinguishedPropertySetIdInternetHeaders DistinguishedPropertySetId = `InternetHeaders`
	// Identifies the meeting property set ID by name.
	DistinguishedPropertySetIdMeeting DistinguishedPropertySetId = `Meeting`
	DistinguishedPropertySetIdSharing DistinguishedPropertySetId = `Sharing`
	// Identifies the public strings property set ID by name.
	DistinguishedPropertySetIdPublicStrings DistinguishedPropertySetId = `PublicStrings`
	// Identifies the task property set ID by name.
	DistinguishedPropertySetIdTask DistinguishedPropertySetId = `Task`
	// Identifies the unified messaging property set ID by name.
	DistinguishedPropertySetIdUnifiedMessaging DistinguishedPropertySetId = `UnifiedMessaging`
)

type PropertyType string

const (
	// A double value that is interpreted as a date and time. The integer part is the date and the fraction part is the time.
	PropertyTypeApplicationTime PropertyType = `PropertyTypeApplicationTime`
	// An array of double values that are interpreted as a date and time.
	PropertyTypeApplicationTimeArray PropertyType = `PropertyTypeApplicationTimeArray`
	// A Base64-encoded binary value.
	PropertyTypeBinary PropertyType = `PropertyTypeBinary`
	// An array of Base64-encoded binary values.
	PropertyTypeBinaryArray PropertyType = `PropertyTypeBinaryArray`
	// A Boolean <strong>true</strong> or <strong>false</strong>.
	PropertyTypeBoolean PropertyType = `PropertyTypeBoolean`
	// A GUID string.
	PropertyTypeCLSID PropertyType = `PropertyTypeCLSID`
	// An array of GUID strings.
	PropertyTypeCLSIDArray PropertyType = `PropertyTypeCLSIDArray`
	// A 64-bit integer that is interpreted as the number of cents.
	PropertyTypeCurrency PropertyType = `PropertyTypeCurrency`
	// An array of 64-bit integers that are interpreted as the number of cents.
	PropertyTypeCurrencyArray PropertyType = `PropertyTypeCurrencyArray`
	// A 64-bit floating-point value.
	PropertyTypeDouble PropertyType = `PropertyTypeDouble`
	// An array of 64-bit floating-point values.
	PropertyTypeDoubleArray PropertyType = `PropertyTypeDoubleArray`
	// SCODE value; 32-bit unsigned integer. Not used for restrictions or for getting/setting values. This exists only for reporting.
	PropertyTypeError PropertyType = `PropertyTypeError`
	// A 32-bit floating-point value.
	PropertyTypeFloat PropertyType = `PropertyTypeFloat`
	// An array of 32-bit floating-point values.
	PropertyTypeFloatArray PropertyType = `PropertyTypeFloatArray`
	// A signed 32-bit (Int32) integer.
	PropertyTypeInteger PropertyType = `PropertyTypeInteger`
	// An array of signed 32-bit (Int32) integers.
	PropertyTypeIntegerArray PropertyType = `PropertyTypeIntegerArray`
	// A signed or unsigned 64-bit (Int64) integer.
	PropertyTypeLong PropertyType = `PropertyTypeLong`
	// An array of signed or unsigned 64-bit (Int64) integers.
	PropertyTypeLongArray PropertyType = `PropertyTypeLongArray`
	// Indicates no property value. Not used for restrictions or for getting/setting values. This exists only for reporting.
	PropertyTypeNull PropertyType = `PropertyTypeNull`
	// A pointer to an object that implements the IUnknown interface. Not used for restrictions or for getting/setting values. This exists only for reporting.
	PropertyTypeObject PropertyType = `PropertyTypeObject`
	// An array of pointers to objects that implement the IUnknown interface. Not used for restrictions or for getting/setting values. This exists only for reporting.
	PropertyTypeObjectArray PropertyType = `PropertyTypeObjectArray`
	// A signed 16-bit integer.
	PropertyTypeShort PropertyType = `PropertyTypeShort`
	// An array of signed 16-bit integers.
	PropertyTypeShortArray PropertyType = `PropertyTypeShortArray`
	// A 64-bit integer data and time value in the form of a FILETIME structure.
	PropertyTypeSystemTime PropertyType = `PropertyTypeSystemTime`
	// An array of 64-bit integer data and time values in the form of a FILETIME structure.
	PropertyTypeSystemTimeArray PropertyType = `PropertyTypeSystemTimeArray`
	// A Unicode string.
	PropertyTypeString PropertyType = `PropertyTypeString`
	// An array of Unicode strings.
	PropertyTypeStringArray PropertyType = `PropertyTypeStringArray`
)
