package elements

// The Urls element specifies an array of URLs for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/urls
import "encoding/xml"

type Urls struct {
	XMLName xml.Name

	// The Url element represents the location of the client Web service for push notifications.
	Url *Url `xml:"Url"`
}

func (U *Urls) SetForMarshal() {
	U.XMLName.Local = "t:Urls"
}

func (U *Urls) GetSchema() *Schema {
	return &SchemaTypes
}
