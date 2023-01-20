package elements

// The DisclaimerText (EncryptionConfigurationResponseType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disclaimertext-encryptionconfigurationresponsetype
import "encoding/xml"

type DisclaimerTextEncryptionConfigurationResponseType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
