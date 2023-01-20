package elements

// The GlobalObjectId element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalobjectid-getclientintenttype
import "encoding/xml"

type GlobalObjectIdGetClientIntentType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
