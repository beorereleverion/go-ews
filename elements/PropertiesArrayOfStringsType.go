package elements

// The Properties element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/properties-arrayofstringstype
import "encoding/xml"

type PropertiesArrayOfStringsType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
