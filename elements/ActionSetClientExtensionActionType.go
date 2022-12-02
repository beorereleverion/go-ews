package elements

// The Action element contains the action that the Exchange server should take on an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/action-setclientextensionactiontype
type ActionSetClientExtensionActionType struct {
	// The ClientExtension element contains user and configuration information about an app.
	ClientExtension *ClientExtension `xml:"t:ClientExtension"`
	// Specifies the identifier of the action. This attribute is required.
	ActionId *string `xml:"ActionId,attr"`
	// Specifies the identifier of the extension. This attribute is optional.
	ExtensionId *string `xml:"ExtensionId,attr"`
}
