package elements

// The GetAppMarketplaceUrl element specifies the request to retrieve the URL for an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmarketplaceurl
import "encoding/xml"

type GetAppMarketplaceUrl struct {
	XMLName xml.Name

	// The ApiVersionSupported element contains the version of the JavaScript API for Office supported by the client.
	ApiVersionSupported *ApiVersionSupported `xml:"ApiVersionSupported"`
	// The SchemaVersionSupported element contains the version of the manifest schema supported by the client.
	SchemaVersionSupported *SchemaVersionSupported `xml:"SchemaVersionSupported"`
}

func (G *GetAppMarketplaceUrl) SetForMarshal() {
	G.XMLName.Local = "m:GetAppMarketplaceUrl"
}

func (G *GetAppMarketplaceUrl) GetSchema() *Schema {
	return &SchemaMessages
}
