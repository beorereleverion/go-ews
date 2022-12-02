package elements

// The SizeRequested element contains the requested photo size for a GetUserPhoto operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sizerequested
type SizeRequested string

const (
	// The image is 120 pixels high and 120 pixels wide.
	SizeRequestedHROneTwoZeroxOneTwoZero SizeRequested = `HR120x120`
	// The image is 240 pixels high and 240 pixels wide.
	SizeRequestedHRTwoFourZeroxTwoFourZero SizeRequested = `HR240x240`
	// The image is 360 pixels high and 360 pixels wide.
	SizeRequestedHRTree6ZeroxTree6Zero SizeRequested = `HR360x360`
	// The image is 432 pixels high and 432 pixels wide.
	SizeRequestedHRFourTreeTwoxFourTreeTwo SizeRequested = `HR432x432`
	// The image is 48 pixels high and 48 pixels wide.
	SizeRequestedHRFour8xFour8 SizeRequested = `HR48x48`
	// The image is 504 pixels high and 504 pixels wide.
	SizeRequestedHR5ZeroFourx5ZeroFour SizeRequested = `HR504x504`
	// The image is 648 pixels high and 648 pixels wide.
	SizeRequestedHR6Four8x6Four8 SizeRequested = `HR648x648`
	// The image is 64 pixels high and 64 pixels wide.
	SizeRequestedHR6Fourx6Four SizeRequested = `HR64x64`
	// The image is 96 pixels high and 96 pixels wide.
	SizeRequestedHR96x96 SizeRequested = `HR96x96`
)
