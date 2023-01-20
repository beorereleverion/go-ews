package elements

// The AppStatus element value indicates the status of the mail app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appstatus-ex15websvcsotherref
import "encoding/xml"

type AppStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The mail app could not be automatically updated. The mail app needs to be re-installed from the Office Store.
	AppStatusOnePointZero string = `1.0`
	// The mail app could not be automatically updated. The mail app requires increased permissions, and this requires your review and confirmation to install.
	AppStatusOnePointOne string = `1.1`
	// The mail app couldn't be updated automatically. The current license has expired or is invalid. Please update the mail app from the Office Store.
	AppStatusOnePointTwo string = `1.2`
	// The mail app license could not be automatically updated. The license for the mail app needs to be recovered from the Office Store.
	AppStatusTwoPointZero string = `2.0`
	// The mail app license could not be automatically updated. The current license has expired. A new license for this app needs to be installed from the Office Store.
	AppStatusTwoPointOne string = `2.1`
	// The Office Store status for the mail app has changed. This may indicate that there is a problem with the mail app. Go to the mail app page in the Office Store for more information.
	AppStatusTreePointZero string = `3.0`
	// The mail app has been removed from the Office Store.
	AppStatusTreePointOne string = `3.1`
	// A problem has been discovered with the mail app and it has temporarily been withdrawn from the Office Store.
	AppStatusTreePointTwo string = `3.2`
	// The mail app will be removed from the Office Store within 30 days.
	AppStatusTreePointTree string = `3.3`
	// The mail app has been automatically disabled by your mail client.
	AppStatusFourPointZero string = `4.0`
	// The mail app has been disabled by Outlook for performance reasons.
	AppStatusFourPointOne string = `4.1`
	// The mail app has a healthy status.
	AppStatusNull string = `Null`
	// The mail app has a healthy status.
	AppStatusZero string = `0`
)

func (A *AppStatus) SetForMarshal() {
	A.XMLName.Local = "t:AppStatus"
}

func (A *AppStatus) GetSchema() *Schema {
	return &SchemaTypes
}
