package elements

// The Urls element specifies an array of URLs that are the result of entity extraction from an item in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/urls-arrayofurlentitiestype
import "encoding/xml"

type UrlsArrayOfUrlEntitiesType struct {
	XMLName xml.Name

	// The UrlEntity element identifies a single extracted URL entity.
	UrlEntity *UrlEntity `xml:"UrlEntity"`
}

func (U *UrlsArrayOfUrlEntitiesType) SetForMarshal() {
	U.XMLName.Local = "t:Urls"
}

func (U *UrlsArrayOfUrlEntitiesType) GetSchema() *Schema {
	return &SchemaTypes
}
