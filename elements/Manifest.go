package elements

// The Manifest element contains the base64-encoded app manifest file.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/manifest
import "encoding/xml"

type Manifest struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
