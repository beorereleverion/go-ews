package elements

// The DisclaimerText (SetEncryptionConfigurationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disclaimertext-setencryptionconfigurationtype
import "encoding/xml"

type DisclaimerTextSetEncryptionConfigurationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
