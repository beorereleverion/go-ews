package elements

// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/to
import "encoding/xml"

type To struct {
	XMLName xml.Name

	// Indicates whether the time zone transition target is a time zone period or of a group of time zone transitions.
	Kind *string `xml:"Kind,attr"`
}

const (
	// Specifies that the time zone transition target is a time zone period.
	ToPeriod = `Period`
	// Specifies that the time zone transition target is a group of time zone transitions.
	ToGroup = `Group`
)

func (T *To) SetForMarshal() {
	T.XMLName.Local = "t:To"
}

func (T *To) GetSchema() *Schema {
	return &SchemaTypes
}
