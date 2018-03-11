package responses

type GeoJson struct {
	Coordinate [][]float64              `json:"coordinate"`
	Properties []map[string]interface{} `json:"properties"`
	Type       string                   `json:"type"`
}
