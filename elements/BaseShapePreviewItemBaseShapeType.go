package elements

// The BaseShape element specifies either the default preview with all properties returned or a compact preview with fewer properties returned.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseshape-previewitembaseshapetype
type BaseShapePreviewItemBaseShapeType string

const (
	// Indicates that only selected properties are shown.
	BaseShapePreviewItemBaseShapeTypeCompact BaseShapePreviewItemBaseShapeType = `Compact`
	// Indicates that all properties are shown.
	BaseShapePreviewItemBaseShapeTypeDefault BaseShapePreviewItemBaseShapeType = `Default`
)
