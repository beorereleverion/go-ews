package elements

// The GetAppManifests element is the base element for a request to return the manifest for apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmanifests
import "encoding/xml"

type GetAppManifests struct {
	XMLName xml.Name

	// The ApiVersionSupported element contains the version of the JavaScript API for Office supported by the client.
	ApiVersionSupported *ApiVersionSupported `xml:"ApiVersionSupported"`
	// The SchemaVersionSupported element contains the version of the manifest schema supported by the client.
	SchemaVersionSupported *SchemaVersionSupported `xml:"SchemaVersionSupported"`
}

func (G *GetAppManifests) SetForMarshal() {
	G.XMLName.Local = "m:GetAppManifests"
}

func (G *GetAppManifests) GetSchema() *Schema {
	return &SchemaMessages
}
