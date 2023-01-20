package elements

// The TransitionsGroups element represents an array of time zone transition groups.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/transitionsgroups
import "encoding/xml"

type TransitionsGroups struct {
	XMLName xml.Name

	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
}

func (T *TransitionsGroups) SetForMarshal() {
	T.XMLName.Local = "t:TransitionsGroups"
}

func (T *TransitionsGroups) GetSchema() *Schema {
	return &SchemaTypes
}
