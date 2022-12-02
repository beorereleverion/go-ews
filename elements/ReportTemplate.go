package elements

// The ReportTemplate element represents the type of report to get.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reporttemplate
type ReportTemplate string

const (
	// RecipientPath
	ReportTemplateRecipientPath ReportTemplate = `RecipientPath`
	// Summary
	ReportTemplateSummary ReportTemplate = `Summary`
)
