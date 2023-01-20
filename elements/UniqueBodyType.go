package elements

// The UniqueBodyType element specifies whether the unique body is returned in text or HTML format.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquebodytype
import "encoding/xml"

type UniqueBodyType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Best
	UniqueBodyTypeBest string = `Best`
	// HTML
	UniqueBodyTypeHTML string = `HTML`
	// Text
	UniqueBodyTypeText string = `Text`
)

func (U *UniqueBodyType) SetForMarshal() {
	U.XMLName.Local = "t:UniqueBodyType"
}

func (U *UniqueBodyType) GetSchema() *Schema {
	return &SchemaTypes
}
