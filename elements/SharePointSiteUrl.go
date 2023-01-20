package elements

// The SharePointSiteURL element contains the Uniform Resource Locator (URL) of the SharePoint site that is linked with the site mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharepointsiteurl
import "encoding/xml"

type SharePointSiteUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SharePointSiteUrl) SetForMarshal() {
	S.XMLName.Local = "m:SharePointSiteUrl"
}

func (S *SharePointSiteUrl) GetSchema() *Schema {
	return &SchemaMessages
}
