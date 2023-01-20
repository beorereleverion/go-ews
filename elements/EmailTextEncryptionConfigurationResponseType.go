package elements

// The EmailText (EncryptionConfigurationResponseType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailtext-encryptionconfigurationresponsetype
import "encoding/xml"

type EmailTextEncryptionConfigurationResponseType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
