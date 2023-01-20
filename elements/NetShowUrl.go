package elements

// The NetShowUrl element specifies the URL for a Microsoft NetShow online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/netshowurl
import "encoding/xml"

type NetShowUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NetShowUrl) SetForMarshal() {
	N.XMLName.Local = "t:NetShowUrl"
}

func (N *NetShowUrl) GetSchema() *Schema {
	return &SchemaTypes
}
