package elements

// The PortalText (EncryptionConfigurationResponseType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/portaltext-encryptionconfigurationresponsetype
import "encoding/xml"

type PortalTextEncryptionConfigurationResponseType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
