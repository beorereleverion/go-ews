package elements

// The ClientExtensions element contains an array of user and configuration information about apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/clientextensions
import "encoding/xml"

type ClientExtensions struct {
	XMLName xml.Name

	// The ClientExtension element contains user and configuration information about an app.
	ClientExtension *ClientExtension `xml:"ClientExtension"`
}
