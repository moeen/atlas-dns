package resources

// DnsResponse represents response from dns route which only has one property and it's the calculated location
type DnsResponse struct {
	Loc float32 `json:"loc"`
}
