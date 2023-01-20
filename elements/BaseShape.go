package elements

// The BaseShape element identifies the set of properties to return in an item or folder response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape
import "encoding/xml"

type BaseShape struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// AllProperties
	BaseShapeAllProperties string = `AllProperties`
	// Default
	BaseShapeDefault string = `Default`
	// IdOnly
	BaseShapeIdOnly string = `IdOnly`
)

func (B *BaseShape) SetForMarshal() {
	B.XMLName.Local = "t:BaseShape"
}

func (B *BaseShape) GetSchema() *Schema {
	return &SchemaTypes
}
