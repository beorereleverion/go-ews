package elements

// The Longitude element specifies the longitude of the location of the associated postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/longitude
import "encoding/xml"

type Longitude struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}
