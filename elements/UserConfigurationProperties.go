package elements

// The UserConfigurationProperties element specifies the property types to get in a GetUserConfiguration operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfigurationproperties
type UserConfigurationProperties string

const (
	// Specifies the identifier, dictionary, XML data, and binary data property types.
	UserConfigurationPropertiesAll UserConfigurationProperties = `All`
	// Specifies binary data property types.
	UserConfigurationPropertiesBinaryData UserConfigurationProperties = `BinaryData`
	// Specifies dictionary property types.
	UserConfigurationPropertiesDictionary UserConfigurationProperties = `Dictionary`
	// Specifies the identifier property.
	UserConfigurationPropertiesId UserConfigurationProperties = `Id`
	// Specifies XML data property types.
	UserConfigurationPropertiesXmlData UserConfigurationProperties = `XmlData`
)
