package elements

// The UrlEntity element identifies a single extracted URL entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/urlentity
import "encoding/xml"

type UrlEntity struct {
	XMLName xml.Name

	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"Position"`
	// The Url element represents the location of the client Web service for push notifications.
	Url *Url `xml:"Url"`
}

func (U *UrlEntity) SetForMarshal() {
	U.XMLName.Local = "t:UrlEntity"
}

func (U *UrlEntity) GetSchema() *Schema {
	return &SchemaTypes
}
