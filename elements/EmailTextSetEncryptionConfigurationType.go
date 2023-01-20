package elements

// The EmailText (SetEncryptionConfigurationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailtext-setencryptionconfigurationtype
import "encoding/xml"

type EmailTextSetEncryptionConfigurationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
