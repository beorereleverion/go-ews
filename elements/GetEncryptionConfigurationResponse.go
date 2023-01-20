package elements

// The GetEncryptionConfigurationResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getencryptionconfigurationresponse
import "encoding/xml"

type GetEncryptionConfigurationResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
