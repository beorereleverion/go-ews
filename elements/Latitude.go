package elements

// The Latitude element specifies the latitude of the location of the associated postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/latitude
import "encoding/xml"

type Latitude struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}
