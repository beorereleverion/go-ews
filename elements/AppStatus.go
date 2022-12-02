package elements

// The AppStatus element value indicates the status of the mail app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appstatus-ex15websvcsotherref
type AppStatus string

const (
	// The mail app could not be automatically updated. The mail app needs to be re-installed from the Office Store.
	AppStatusOnePointZero AppStatus = `1.0`
	// The mail app could not be automatically updated. The mail app requires increased permissions, and this requires your review and confirmation to install.
	AppStatusOnePointOne AppStatus = `1.1`
	// The mail app couldn't be updated automatically. The current license has expired or is invalid. Please update the mail app from the Office Store.
	AppStatusOnePointTwo AppStatus = `1.2`
	// The mail app license could not be automatically updated. The license for the mail app needs to be recovered from the Office Store.
	AppStatusTwoPointZero AppStatus = `2.0`
	// The mail app license could not be automatically updated. The current license has expired. A new license for this app needs to be installed from the Office Store.
	AppStatusTwoPointOne AppStatus = `2.1`
	// The Office Store status for the mail app has changed. This may indicate that there is a problem with the mail app. Go to the mail app page in the Office Store for more information.
	AppStatusTreePointZero AppStatus = `3.0`
	// The mail app has been removed from the Office Store.
	AppStatusTreePointOne AppStatus = `3.1`
	// A problem has been discovered with the mail app and it has temporarily been withdrawn from the Office Store.
	AppStatusTreePointTwo AppStatus = `3.2`
	// The mail app will be removed from the Office Store within 30 days.
	AppStatusTreePointTree AppStatus = `3.3`
	// The mail app has been automatically disabled by your mail client.
	AppStatusFourPointZero AppStatus = `4.0`
	// The mail app has been disabled by Outlook for performance reasons.
	AppStatusFourPointOne AppStatus = `4.1`
	// The mail app has a healthy status.
	AppStatusNull AppStatus = `Null`
	// The mail app has a healthy status.
	AppStatusZero AppStatus = `0`
)
