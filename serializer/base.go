package serializer

import (
	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/responses"
)

func NewPlace(pb *pbnavitia.PtObject) *responses.Place {
	if pb == nil {
		return nil
	}
	t := pb.EmbeddedType.String()
	place := responses.Place{
		Id:           pb.Uri,
		Name:         pb.Name,
		EmbeddedType: &t,
		Quality:      pb.Quality,
	}
	return &place
}
