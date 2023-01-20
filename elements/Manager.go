package elements

// The Manager element represents a contact's manager.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/manager
import "encoding/xml"

type Manager struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *Manager) SetForMarshal() {
	M.XMLName.Local = "t:Manager"
}

func (M *Manager) GetSchema() *Schema {
	return &SchemaTypes
}
