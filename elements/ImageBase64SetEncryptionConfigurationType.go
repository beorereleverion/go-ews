package elements

// The ImageBase64 (SetEncryptionConfigurationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imagebase64-setencryptionconfigurationtype
import "encoding/xml"

type ImageBase64SetEncryptionConfigurationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
