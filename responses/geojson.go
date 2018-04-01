package responses

type GeoJson struct {
	Coordinates [][]float64              `json:"coordinates"`
	Properties  []map[string]interface{} `json:"properties,omitempty"`
	Type        string                   `json:"type"`
}

type GeoJsonMultilineString struct {
	Coordinates [][][]float64            `json:"coordinates"`
	Properties  []map[string]interface{} `json:"properties,omitempty"`
	Type        string                   `json:"type"`
}
