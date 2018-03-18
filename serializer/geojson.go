package serializer

import "github.com/kinnou02/gonavitia/responses"
import "github.com/kinnou02/gonavitia/pbnavitia"

func NewGeoJson(pb *pbnavitia.Section) *responses.GeoJson {
	if pb == nil {
		return nil
	}
	g := responses.GeoJson{
		Type:        "LineString",
		Properties:  []map[string]interface{}{{"length": pb.GetLength()}},
		Coordinates: make([][]float64, 0, len(pb.Shape)),
	}

	for _, pb_coord := range pb.Shape {
		coord := []float64{pb_coord.GetLon(), pb_coord.GetLat()}
		g.Coordinates = append(g.Coordinates, coord)
	}

	return &g
}
