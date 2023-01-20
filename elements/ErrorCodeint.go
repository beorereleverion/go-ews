package elements

// The ErrorCode element specifies the error code of a failed search performed against a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errorcode-int
import "encoding/xml"

type ErrorCodeint struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (E *ErrorCodeint) SetForMarshal() {
	E.XMLName.Local = "t:ErrorCode"
}

func (E *ErrorCodeint) GetSchema() *Schema {
	return &SchemaTypes
}
