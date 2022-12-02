package elements

// The ClientExtension element contains user and configuration information about an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/clientextension
type ClientExtension struct {
	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
	// The SpecificUsers element specifies the email accounts that can access the app.
	SpecificUsers *SpecificUsers `xml:"t:SpecificUsers"`
	// Specifies the status code of a mail app in an unexpected state.
	AppStatus *string `xml:"AppStatus,attr"`
	// Specifies the license token for paid or trial mail apps.
	Etoken *string `xml:"Etoken,attr"`
	// Specifies whether the app is available. A text value of true for the IsAvailable attribute indicates that the app is available. A value of false indicates that the app is not available. This attribute is optional.
	IsAvailable *string `xml:"IsAvailable,attr"`
	// Specifies whether the app is enabled by default. A text value of true for the IsEnabledByDefault attribute indicates that the app is enabled by default. A value of false indicates that the app is not enabled by default. This attribute is optional.
	IsEnabledByDefault *string `xml:"IsEnabledByDefault,attr"`
	// Specifies whether the app is mandatory. A text value of true for the IsMandatory attribute indicates that the app is mandatory for the mailbox. A value of false indicates that the app is not mandatory. This attribute is optional.
	IsMandatory *string `xml:"IsMandatory,attr"`
	// Specifies the marketplace asset identifier of the app.
	MarketplaceAssetId *string `xml:"MarketplaceAssetId,attr"`
	// Specifies the marketplace content that a user sees for details and reviews about an app.
	MarketplaceContentMarket *string `xml:"MarketplaceContentMarket,attr"`
	// Specifies to whom the app is provided. This attribute is optional.
	ProvidedTo *string `xml:"ProvidedTo,attr"`
	// Specifies the scope of the app.
	Scope *string `xml:"Scope,attr"`
	// Specifies the type of the app.
	Type *string `xml:"Type,attr"`
}
