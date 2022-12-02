package elements

// The Type element specifies a dictionary object type.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/type-userconfiguration
type TypeUserConfiguration string

const (
	// Boolean
	TypeUserConfigurationBoolean TypeUserConfiguration = `Boolean`
	// Byte
	TypeUserConfigurationByte TypeUserConfiguration = `Byte`
	// ByteArray
	TypeUserConfigurationByteArray TypeUserConfiguration = `ByteArray`
	// DateTime
	TypeUserConfigurationDateTime TypeUserConfiguration = `DateTime`
	// Integer32
	TypeUserConfigurationIntegerTreeTwo TypeUserConfiguration = `Integer32`
	// Integer64
	TypeUserConfigurationInteger6Four TypeUserConfiguration = `Integer64`
	// String
	TypeUserConfigurationString TypeUserConfiguration = `String`
	// StringArray
	TypeUserConfigurationStringArray TypeUserConfiguration = `StringArray`
	// UnsignedInteger32
	TypeUserConfigurationUnsignedIntegerTreeTwo TypeUserConfiguration = `UnsignedInteger32`
	// UnsignedInteger64
	TypeUserConfigurationUnsignedInteger6Four TypeUserConfiguration = `UnsignedInteger64`
)
