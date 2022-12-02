package elements

// The RequestServerVersion element contains the versioning information that identifies the schema version to target for a request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestserverversion
type RequestServerVersion struct {
	// Describes the version to target for the request. This attribute is required when the target server version is a version of Exchange starting with Exchange Server 2010.
	Version *string `xml:"Version,attr"`
}

const (
	// Target the schema files for the initial release version of Exchange 2007.
	RequestServerVersionExchange2007 = `Exchange2007`
	// Target the schema files for Exchange 2007 Service Pack 1 (SP1), Exchange 2007 Service Pack 2 (SP2), and Exchange 2007 Service Pack 3 (SP3).
	RequestServerVersionExchange2007_SP1 = `Exchange2007_SP1`
	// Target the schema files for Exchange 2010.
	RequestServerVersionExchange2010 = `Exchange2010`
	// Target the schema files for Exchange 2010 Service Pack 1 (SP1).
	RequestServerVersionExchange2010_SP1 = `Exchange2010_SP1`
	// Target the schema files for Exchange 2010 Service Pack 2 (SP2) and Exchange 2010 Service Pack 3 (SP3).
	RequestServerVersionExchange2010_SP2 = `Exchange2010_SP2`
	// Target the schema files for Exchange 2013.
	RequestServerVersionExchange2013 = `Exchange2013`
	// Target the schema files for Exchange 2013 Service Pack 1 (SP1).
	RequestServerVersionExchange2013_SP1 = `Exchange2013_SP1`
)
