package shodan

type HostLocation struct {
	City         string  `json:"region_code"`
	RegionCOde   string  `json:"area_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code`
	Latitude     float32 `json:"latitude"`
}

type Host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	HostNames []string     `json:hostnames`
	Location  HostLocation `json:"location"  `
}

	type HostSearch struct {
	Matches []Host `json:"matches"`
}
