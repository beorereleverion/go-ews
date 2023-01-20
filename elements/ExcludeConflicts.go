package elements

// The ExcludeConflicts element specifies whether to return suggested times for calendar times that conflict among the attendees.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/excludeconflicts
import "encoding/xml"

type ExcludeConflicts struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ExcludeConflictsfalse bool = false
	// true
	ExcludeConflictstrue bool = true
)

func (E *ExcludeConflicts) SetForMarshal() {
	E.XMLName.Local = "t:ExcludeConflicts"
}

func (E *ExcludeConflicts) GetSchema() *Schema {
	return &SchemaTypes
}
