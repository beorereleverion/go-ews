package elements

// The IsPermanentFailure element indicates whether a previous attempt to index the item was unsuccessful.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispermanentfailure
import "encoding/xml"

type IsPermanentFailure struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsPermanentFailurefalse bool = false
	// true
	IsPermanentFailuretrue bool = true
)

func (I *IsPermanentFailure) SetForMarshal() {
	I.XMLName.Local = "t:IsPermanentFailure"
}

func (I *IsPermanentFailure) GetSchema() *Schema {
	return &SchemaTypes
}
