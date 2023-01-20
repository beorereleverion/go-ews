package elements

// The AppMarketplaceUrl element specifies the URL for the app marketplace.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appmarketplaceurl
import "encoding/xml"

type AppMarketplaceUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *AppMarketplaceUrl) SetForMarshal() {
	A.XMLName.Local = "m:AppMarketplaceUrl"
}

func (A *AppMarketplaceUrl) GetSchema() *Schema {
	return &SchemaMessages
}
