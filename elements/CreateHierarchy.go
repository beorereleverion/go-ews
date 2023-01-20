package elements

// The CreateHierarchy element indicates whether a client can create a hierarchy table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createhierarchy
import "encoding/xml"

type CreateHierarchy struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	CreateHierarchyfalse bool = false
	// true
	CreateHierarchytrue bool = true
)

func (C *CreateHierarchy) SetForMarshal() {
	C.XMLName.Local = "t:CreateHierarchy"
}

func (C *CreateHierarchy) GetSchema() *Schema {
	return &SchemaTypes
}
