package elements

// The ImageBase64 element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imagebase64-encryptionconfigurationresponsetype
import "encoding/xml"

type ImageBase64EncryptionConfigurationResponseType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
