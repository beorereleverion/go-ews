package elements

// The DiagnosticsLevel element represents timing and performance information that will be used to derive the report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/diagnosticslevel
import "encoding/xml"

type DiagnosticsLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DiagnosticsLevel) SetForMarshal() {
	D.XMLName.Local = "m:DiagnosticsLevel"
}

func (D *DiagnosticsLevel) GetSchema() *Schema {
	return &SchemaMessages
}
