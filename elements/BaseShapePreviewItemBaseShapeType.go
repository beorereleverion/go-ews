package elements

// The BaseShape element specifies either the default preview with all properties returned or a compact preview with fewer properties returned.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape-previewitembaseshapetype
import "encoding/xml"

type BaseShapePreviewItemBaseShapeType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Indicates that only selected properties are shown.
	BaseShapePreviewItemBaseShapeTypeCompact string = `Compact`
	// Indicates that all properties are shown.
	BaseShapePreviewItemBaseShapeTypeDefault string = `Default`
)

func (B *BaseShapePreviewItemBaseShapeType) SetForMarshal() {
	B.XMLName.Local = "t:BaseShape"
}

func (B *BaseShapePreviewItemBaseShapeType) GetSchema() *Schema {
	return &SchemaTypes
}
