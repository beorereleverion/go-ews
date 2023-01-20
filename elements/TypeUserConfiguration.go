package elements

// The Type element specifies a dictionary object type.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/type-userconfiguration
import "encoding/xml"

type TypeUserConfiguration struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Boolean
	TypeUserConfigurationBoolean string = `Boolean`
	// Byte
	TypeUserConfigurationByte string = `Byte`
	// ByteArray
	TypeUserConfigurationByteArray string = `ByteArray`
	// DateTime
	TypeUserConfigurationDateTime string = `DateTime`
	// Integer32
	TypeUserConfigurationIntegerTreeTwo string = `Integer32`
	// Integer64
	TypeUserConfigurationInteger6Four string = `Integer64`
	// String
	TypeUserConfigurationString string = `String`
	// StringArray
	TypeUserConfigurationStringArray string = `StringArray`
	// UnsignedInteger32
	TypeUserConfigurationUnsignedIntegerTreeTwo string = `UnsignedInteger32`
	// UnsignedInteger64
	TypeUserConfigurationUnsignedInteger6Four string = `UnsignedInteger64`
)

func (T *TypeUserConfiguration) SetForMarshal() {
	T.XMLName.Local = "t:Type"
}

func (T *TypeUserConfiguration) GetSchema() *Schema {
	return &SchemaTypes
}
