package elements

// The GetEncryptionConfigurationResponseMessage is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getencryptionconfigurationresponsemessage
import "encoding/xml"

type GetEncryptionConfigurationResponseMessage struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
