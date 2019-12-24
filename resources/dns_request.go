package resources

// DnsRequest represents a request body coming from dns route
// If we don't use pointers for this struct members, validator will consider a DnsRequest with zero values as an invalid
// data since they are required
type DnsRequest struct {
	X   *float32 `json:"x,string" validate:"required"`
	Y   *float32 `json:"y,string" validate:"required"`
	Z   *float32 `json:"z,string" validate:"required"`
	Vel *float32 `json:"vel,string" validate:"required"`
}
