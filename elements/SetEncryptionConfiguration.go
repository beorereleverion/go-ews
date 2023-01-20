package elements

// The SetEncryptionConfiguration element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setencryptionconfiguration
import "encoding/xml"

type SetEncryptionConfiguration struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
