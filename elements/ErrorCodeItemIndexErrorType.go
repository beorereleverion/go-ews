package elements

// The ErrorCode element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errorcode-itemindexerrortype
import "encoding/xml"

type ErrorCodeItemIndexErrorType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (E *ErrorCodeItemIndexErrorType) SetForMarshal() {
	E.XMLName.Local = "t:ErrorCode"
}

func (E *ErrorCodeItemIndexErrorType) GetSchema() *Schema {
	return &SchemaTypes
}
