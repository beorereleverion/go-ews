package elements

// The EmailAddress element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddress-emailaddress
import "encoding/xml"

type EmailAddressEmailAddress struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
