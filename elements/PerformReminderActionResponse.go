package elements

// The PerformReminderActionResponse element specifies the response to a PerformReminderAction request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/performreminderactionresponse
type PerformReminderActionResponse struct {
	// The UpdatedItemIds element specifies the identifiers of updated reminder items.
	UpdatedItemIds *UpdatedItemIds `xml:"m:UpdatedItemIds"`
}
