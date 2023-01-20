package elements

// The SetEncryptionConfigurationResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setencryptionconfigurationresponse
import "encoding/xml"

type SetEncryptionConfigurationResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
