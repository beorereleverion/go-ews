package elements

// The ApiVersionSupported element contains the version of the JavaScript API for Office supported by the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/apiversionsupported
import "encoding/xml"

type ApiVersionSupported struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *ApiVersionSupported) SetForMarshal() {
	A.XMLName.Local = "m:ApiVersionSupported"
}

func (A *ApiVersionSupported) GetSchema() *Schema {
	return &SchemaMessages
}
