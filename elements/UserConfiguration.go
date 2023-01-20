package elements

// The UserConfiguration element defines a single user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfiguration
import "encoding/xml"

type UserConfiguration struct {
	XMLName xml.Name

	// The BinaryData element contains binary data property content.
	BinaryData *BinaryData `xml:"BinaryData"`
	// The Dictionary element defines a set of dictionary property entries for a user configuration object.
	Dictionary *Dictionary `xml:"Dictionary"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
	UserConfigurationName *UserConfigurationName `xml:"UserConfigurationName"`
	// The XmlData element contains XML data property content for a user configuration object.
	XmlData *XmlData `xml:"XmlData"`
}

func (U *UserConfiguration) SetForMarshal() {
	U.XMLName.Local = "m:UserConfiguration"
}

func (U *UserConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
