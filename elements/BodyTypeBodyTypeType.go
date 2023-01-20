package elements

// The BodyType element specifies the type of the body of the item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodytype-bodytypetype
import "encoding/xml"

type BodyTypeBodyTypeType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Indicates that the body is HTML.
	BodyTypeBodyTypeTypeHTML string = `HTML`
	// Indicates that the body is text.
	BodyTypeBodyTypeTypeText string = `Text`
)

func (B *BodyTypeBodyTypeType) SetForMarshal() {
	B.XMLName.Local = "t:BodyType"
}

func (B *BodyTypeBodyTypeType) GetSchema() *Schema {
	return &SchemaTypes
}
