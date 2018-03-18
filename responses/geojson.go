package responses

type GeoJson struct {
	Coordinates [][]float64              `json:"coordinates"`
	Properties  []map[string]interface{} `json:"properties"`
	Type        string                   `json:"type"`
}
