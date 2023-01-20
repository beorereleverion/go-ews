package elements

// The UserConfigurationProperties element specifies the property types to get in a GetUserConfiguration operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfigurationproperties
import "encoding/xml"

type UserConfigurationProperties struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Specifies the identifier, dictionary, XML data, and binary data property types.
	UserConfigurationPropertiesAll string = `All`
	// Specifies binary data property types.
	UserConfigurationPropertiesBinaryData string = `BinaryData`
	// Specifies dictionary property types.
	UserConfigurationPropertiesDictionary string = `Dictionary`
	// Specifies the identifier property.
	UserConfigurationPropertiesId string = `Id`
	// Specifies XML data property types.
	UserConfigurationPropertiesXmlData string = `XmlData`
)

func (U *UserConfigurationProperties) SetForMarshal() {
	U.XMLName.Local = "m:UserConfigurationProperties"
}

func (U *UserConfigurationProperties) GetSchema() *Schema {
	return &SchemaMessages
}
