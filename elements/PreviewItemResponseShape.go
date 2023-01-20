package elements

// The PreviewItemResponseShape element contains the requested property set to be returned in a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/previewitemresponseshape
import "encoding/xml"

type PreviewItemResponseShape struct {
	XMLName xml.Name

	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"AdditionalProperties"`
	// The BaseShape element specifies either the default preview with all properties returned or a compact preview with fewer properties returned.
	BaseShape *BaseShapePreviewItemBaseShapeType `xml:"BaseShape"`
}

func (P *PreviewItemResponseShape) SetForMarshal() {
	P.XMLName.Local = "m:PreviewItemResponseShape"
}

func (P *PreviewItemResponseShape) GetSchema() *Schema {
	return &SchemaMessages
}
