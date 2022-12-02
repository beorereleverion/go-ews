package elements

// The UserConfiguration element defines a single user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfiguration
type UserConfiguration struct {
	// The BinaryData element contains binary data property content.
	BinaryData *BinaryData `xml:"t:BinaryData"`
	// The Dictionary element defines a set of dictionary property entries for a user configuration object.
	Dictionary *Dictionary `xml:"t:Dictionary"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
	UserConfigurationName *UserConfigurationName `xml:"t:UserConfigurationName"`
	// The XmlData element contains XML data property content for a user configuration object.
	XmlData *XmlData `xml:"t:XmlData"`
}
