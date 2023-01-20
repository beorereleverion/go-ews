package elements

// The ChangeCount element specifies the version of the task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changecount
import "encoding/xml"

type ChangeCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *ChangeCount) SetForMarshal() {
	C.XMLName.Local = "t:ChangeCount"
}

func (C *ChangeCount) GetSchema() *Schema {
	return &SchemaTypes
}
