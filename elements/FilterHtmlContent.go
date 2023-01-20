package elements

// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/filterhtmlcontent
import "encoding/xml"

type FilterHtmlContent struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	FilterHtmlContentfalse bool = false
	// true
	FilterHtmlContenttrue bool = true
)

func (F *FilterHtmlContent) SetForMarshal() {
	F.XMLName.Local = "t:FilterHtmlContent"
}

func (F *FilterHtmlContent) GetSchema() *Schema {
	return &SchemaTypes
}
