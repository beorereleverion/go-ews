package elements

// The Timeout element represents the duration, in minutes, that the subscription can remain idle without a GetEvents request from the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timeout
type Timeout int64
