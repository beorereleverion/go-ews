package elements

// The ContentExpiryDate element specifies the expiration date of the content.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contentexpirydate
import "encoding/xml"

type ContentExpiryDate struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContentExpiryDate) SetForMarshal() {
	C.XMLName.Local = "t:ContentExpiryDate"
}

func (C *ContentExpiryDate) GetSchema() *Schema {
	return &SchemaTypes
}
