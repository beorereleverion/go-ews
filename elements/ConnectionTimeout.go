package elements

// The ConnectionTimeout element specifies the number of minutes to keep a connection open.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectiontimeout
import "encoding/xml"

type ConnectionTimeout struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *ConnectionTimeout) SetForMarshal() {
	C.XMLName.Local = "t:ConnectionTimeout"
}

func (C *ConnectionTimeout) GetSchema() *Schema {
	return &SchemaTypes
}
