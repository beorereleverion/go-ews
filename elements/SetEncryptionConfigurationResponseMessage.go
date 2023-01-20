package elements

// The SetEncryptionConfigurationResponseMessage element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setencryptionconfigurationresponsemessage
import "encoding/xml"

type SetEncryptionConfigurationResponseMessage struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
