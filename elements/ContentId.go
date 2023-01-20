package elements

// The ContentId element represents an identifier for the contents of an attachment. ContentId can be set to any string value. Applications can use ContentId to implement their own identification mechanisms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contentid
import "encoding/xml"

type ContentId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContentId) SetForMarshal() {
	C.XMLName.Local = "t:ContentId"
}

func (C *ContentId) GetSchema() *Schema {
	return &SchemaTypes
}
