package elements

// The GetEncryptionConfiguration element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getencryptionconfiguration
import "encoding/xml"

type GetEncryptionConfiguration struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
