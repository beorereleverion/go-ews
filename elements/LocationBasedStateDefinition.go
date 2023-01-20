package elements

// The LocationBasedStateDefinition element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/locationbasedstatedefinition
import "encoding/xml"

type LocationBasedStateDefinition struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
