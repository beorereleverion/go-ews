package elements

// The IdentityDetails element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/identitydetails
import "encoding/xml"

type IdentityDetails struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
