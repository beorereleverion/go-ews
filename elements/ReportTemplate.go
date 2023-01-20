package elements

// The ReportTemplate element represents the type of report to get.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reporttemplate
import "encoding/xml"

type ReportTemplate struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// RecipientPath
	ReportTemplateRecipientPath string = `RecipientPath`
	// Summary
	ReportTemplateSummary string = `Summary`
)

func (R *ReportTemplate) SetForMarshal() {
	R.XMLName.Local = "m:ReportTemplate"
}

func (R *ReportTemplate) GetSchema() *Schema {
	return &SchemaMessages
}
