package elements

// The NormalizedBodyType element specifies whether the normalized body is returned in text or HTML format.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/normalizedbodytype
import "encoding/xml"

type NormalizedBodyType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Best
	NormalizedBodyTypeBest string = `Best`
	// HTML
	NormalizedBodyTypeHTML string = `HTML`
	// Text
	NormalizedBodyTypeText string = `Text`
)

func (N *NormalizedBodyType) SetForMarshal() {
	N.XMLName.Local = "t:NormalizedBodyType"
}

func (N *NormalizedBodyType) GetSchema() *Schema {
	return &SchemaTypes
}
