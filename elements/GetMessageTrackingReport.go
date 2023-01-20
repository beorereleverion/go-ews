package elements

// The GetMessageTrackingReport element contains the request for the GetMessageTrackingReport operation to retrieve the full message tracking report for the specified ID.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getmessagetrackingreport
import "encoding/xml"

type GetMessageTrackingReport struct {
	XMLName xml.Name

	// The DiagnosticsLevel element represents timing and performance information that will be used to derive the report.
	DiagnosticsLevel *DiagnosticsLevel `xml:"DiagnosticsLevel"`
	// The MessageTrackingReportId element represents the message by its message ID, the organization where the message was found, the server on which the message was submitted, and an internal ID that uniquely identifies the message.
	MessageTrackingReportId *MessageTrackingReportId `xml:"MessageTrackingReportId"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The RecipientFilter element represents a recipient address to use with the specified message tracking report.
	RecipientFilter *RecipientFilter `xml:"RecipientFilter"`
	// The ReportTemplate element represents the type of report to get.
	ReportTemplate *ReportTemplate `xml:"ReportTemplate"`
	// The ReturnQueueEvents element indicates that the person who is running the task is in a privileged role.
	ReturnQueueEvents *ReturnQueueEvents `xml:"ReturnQueueEvents"`
	// The Scope element specifies the scope of the message tracking report.
	Scope *ScopeNonEmptyStringType `xml:"Scope"`
}

func (G *GetMessageTrackingReport) SetForMarshal() {
	G.XMLName.Local = "m:GetMessageTrackingReport"
}

func (G *GetMessageTrackingReport) GetSchema() *Schema {
	return &SchemaMessages
}
