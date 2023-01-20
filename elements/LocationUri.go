package elements

// The LocationUri element contains a string specifying a Uniform Resource Identifier (URI) of the associated postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/locationuri
import "encoding/xml"

type LocationUri struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
