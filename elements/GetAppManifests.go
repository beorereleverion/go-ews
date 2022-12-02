package elements

// The GetAppManifests element is the base element for a request to return the manifest for apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmanifests
type GetAppManifests struct {
	// The ApiVersionSupported element contains the version of the JavaScript API for Office supported by the client.
	ApiVersionSupported *ApiVersionSupported `xml:"m:ApiVersionSupported"`
	// The SchemaVersionSupported element contains the version of the manifest schema supported by the client.
	SchemaVersionSupported *SchemaVersionSupported `xml:"m:SchemaVersionSupported"`
}
