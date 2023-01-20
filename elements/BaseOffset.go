package elements

// The BaseOffset element represents the hourly offset from Coordinated Universal Time (UTC) for the current time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/baseoffset
import "encoding/xml"

type BaseOffset struct {
	XMLName xml.Name
	//this is string, because BaseOffset has strange format
	TEXT *string `xml:",chardata"`
}

func (B *BaseOffset) SetForMarshal() {
	B.XMLName.Local = "t:BaseOffset"
}

func (B *BaseOffset) GetSchema() *Schema {
	return &SchemaTypes
}
