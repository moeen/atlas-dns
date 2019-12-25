package resources

// DnsRequest represents a request body coming from dns route
// If we don't use pointers for this struct members, validator will consider a DnsRequest with zero values as an invalid
// data since they are required
type DnsRequest struct {
	X   *float32 `json:"x,string" binding:"required"`
	Y   *float32 `json:"y,string" binding:"required"`
	Z   *float32 `json:"z,string" binding:"required"`
	Vel *float32 `json:"vel,string" binding:"required"`
}
