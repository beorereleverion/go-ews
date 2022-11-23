package operations

type Schema string

const (
	SchemaSOAP     Schema = `http://schemas.xmlsoap.org/soap/envelope/`
	SchemaTypes    Schema = `http://schemas.microsoft.com/exchange/services/2006/types`
	SchemaMessages Schema = `http://schemas.microsoft.com/exchange/services/2006/messages`
)
